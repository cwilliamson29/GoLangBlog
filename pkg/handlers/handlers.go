package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/models"
	"github.com/cwilliamson29/GoLangBlog/pkg/config"
	"github.com/cwilliamson29/GoLangBlog/pkg/dbdriver"
	"github.com/cwilliamson29/GoLangBlog/pkg/forms"
	"github.com/cwilliamson29/GoLangBlog/pkg/repository"
	"github.com/cwilliamson29/GoLangBlog/pkg/repository/dbrepo"
	"github.com/justinas/nosurf"
	"log"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
	//AdminTemplates *template.Template
	//UITemplates    *template.Template
}

var Repo *Repository

func NewRepo(ac *config.AppConfig, db *dbdriver.DB) *Repository {
	return &Repository{
		App: ac,
		DB:  dbrepo.NewPostGresRepo(db.SQL, ac),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) AddCSRFData(pd *models.PageData, r *http.Request) *models.PageData {
	pd.CSRFToken = nosurf.Token(r)

	if m.App.Session.Exists(r.Context(), "user_id") {
		pd.IsAuthenticated = 1
	}
	return pd
}

// HomeHandler - for getting the home page
func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	//ut := m.App.Session.Get(r.Context(), "user_type")
	//log.Println("user type: ", ut)

	pd := m.AddCSRFData(&models.PageData{}, r)

	var artList map[int]interface{}
	artList, err := m.DB.Get3BlogPost()
	if err != nil {
		log.Println(err)
		return
	}

	err = m.App.UITemplates.ExecuteTemplate(w, "home.page.tmpl", &models.PageData{
		Data:            artList,
		CSRFToken:       pd.CSRFToken,
		IsAuthenticated: pd.IsAuthenticated,
	})
	if err != nil {
		log.Println(err)
	}
}

// AboutHandler - for getting the about page
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	//strMap := make(map[string]string)
	//render.RenderTemplate(w, r, "about.page.tmpl", &models.PageData{StrMap: strMap})
	pd := m.AddCSRFData(&models.PageData{}, r)

	err := m.App.UITemplates.ExecuteTemplate(w, "about.page.tmpl", &models.PageData{
		CSRFToken:       pd.CSRFToken,
		IsAuthenticated: pd.IsAuthenticated})
	if err != nil {
		return
	}
}

// LoginHandler - for getting the login page
func (m *Repository) LoginHandler(w http.ResponseWriter, r *http.Request) {
	pd := m.AddCSRFData(&models.PageData{}, r)

	err := m.App.UITemplates.ExecuteTemplate(w, "login.page.tmpl", &models.PageData{
		CSRFToken:       pd.CSRFToken,
		IsAuthenticated: pd.IsAuthenticated})
	if err != nil {
		http.Error(w, "unable to execute the template", http.StatusInternalServerError)
		return
	}
}

// PageHandler - for getting the individual pages
func (m *Repository) PageHandler(w http.ResponseWriter, r *http.Request) {
	err := m.App.UITemplates.ExecuteTemplate(w, "page.page.tmpl", &models.PageData{})
	if err != nil {
		return
	}

}

// MakePostHandler - for creating new posts
func (m *Repository) MakePostHandler(w http.ResponseWriter, r *http.Request) {
	pd := m.AddCSRFData(&models.PageData{}, r)

	if !m.App.Session.Exists(r.Context(), "user_id") {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	}
	var emptyArticle models.Article
	data := make(map[string]interface{})
	data["article"] = emptyArticle

	//render.RenderTemplate(w, r, "make-post.page.tmpl", &models.PageData{
	//	Form:            forms.New(nil),
	//	Data:            data,
	//	CSRFToken:       pd.CSRFToken,
	//	IsAuthenticated: pd.IsAuthenticated,
	//})
	err := m.App.UITemplates.ExecuteTemplate(w, "make-post.page.tmpl", &models.PageData{
		Form: forms.New(nil),
		//Data:            data,
		CSRFToken:       pd.CSRFToken,
		IsAuthenticated: pd.IsAuthenticated,
	})
	if err != nil {
		return
	}

}

// PostMakePostHandler - Post method for submitting new posts
func (m *Repository) PostMakePostHandler(w http.ResponseWriter, r *http.Request) {
	pd := m.AddCSRFData(&models.PageData{}, r)

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	uID := (m.App.Session.Get(r.Context(), "user_id")).(int)
	article := models.Post{
		Title:   r.Form.Get("blog_title"),
		Content: r.Form.Get("blog_article"),
		UserID:  int(uID),
	}

	form := forms.New(r.PostForm)

	form.HasRequired("blog_title", "blog_article")

	form.MinLength("blog_title", 5, r)
	form.MinLength("blog_article", 5, r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["article"] = article

		//render.RenderTemplate(w, r, "make-post.page.tmpl", &models.PageData{
		//	Form: form,
		//	Data: data,
		//})
		err := m.App.UITemplates.ExecuteTemplate(w, "make-post.page.tmpl", &models.PageData{
			Form: form,
			//Data:            data,
			CSRFToken:       pd.CSRFToken,
			IsAuthenticated: pd.IsAuthenticated,
		})
		if err != nil {
			return
		}
		return
	}
	// Write to the DB
	err = m.DB.InsertPost(article)
	if err != nil {
		log.Fatal(err)
	}

	m.App.Session.Put(r.Context(), "article", article)
	http.Redirect(w, r, "/article-received", http.StatusSeeOther)
}

// ArticleReceived - get article
func (m *Repository) ArticleReceived(w http.ResponseWriter, r *http.Request) {
	pd := m.AddCSRFData(&models.PageData{}, r)

	article, ok := m.App.Session.Get(r.Context(), "article").(models.Article)
	if !ok {
		log.Println("Cant get data from session")
		m.App.Session.Put(r.Context(), "error", "Cant get data from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	data := make(map[string]interface{})
	data["article"] = article

	//render.RenderTemplate(w, r, "article-received.page.tmpl", &models.PageData{
	//	Data: data,
	//})
	err := m.App.UITemplates.ExecuteTemplate(w, "article-received.page.tmpl", &models.PageData{
		//Data:            data,
		CSRFToken:       pd.CSRFToken,
		IsAuthenticated: pd.IsAuthenticated,
	})
	if err != nil {
		return
	}
}

// PostLoginHandler - for getting the individual pages
func (m *Repository) PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.RenewToken(r.Context())
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	form := forms.New(r.PostForm)
	form.HasRequired("email", "password")
	form.IsEmail("email")

	if !form.Valid() {
		err := m.App.UITemplates.ExecuteTemplate(w, "login.page.tmpl", &models.PageData{Form: form})
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	id, _, err := m.DB.AuthenticateUser(email, password)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Invalid Email OR Password")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	m.App.Session.Put(r.Context(), "user_id", id)
	m.App.Session.Put(r.Context(), "flash", "Valid Login")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (m *Repository) PostAdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.RenewToken(r.Context())
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	form := forms.New(r.PostForm)
	form.HasRequired("email", "password")
	form.IsEmail("email")

	if !form.Valid() {
		err := m.App.UITemplates.ExecuteTemplate(w, "login.page.tmpl", &models.PageData{Form: form})
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	id, _, err := m.DB.AuthenticateUser(email, password)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "Invalid Email OR Password")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}
	m.App.Session.Put(r.Context(), "user_id", id)
	m.App.Session.Put(r.Context(), "flash", "Valid Login")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (m *Repository) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.Destroy(r.Context())
	_ = m.App.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
