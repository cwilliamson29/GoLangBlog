package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/models"
	"github.com/cwilliamson29/GoLangBlog/pkg/config"
	"github.com/cwilliamson29/GoLangBlog/pkg/dbRepo"
	"github.com/cwilliamson29/GoLangBlog/pkg/dbdriver"
	"github.com/cwilliamson29/GoLangBlog/pkg/forms"
	"github.com/justinas/nosurf"
	"log"
	"net/http"
)

type BHandlers struct {
	App *config.AppConfig
	DB  dbRepo.DatabaseRepo
}

var Repo *BHandlers

func NewRepo(ac *config.AppConfig, db *dbdriver.DB) *BHandlers {
	return &BHandlers{
		App: ac,
		DB:  dbRepo.NewSQLRepo(db.SQL, ac),
	}
}

func NewHandlers(r *BHandlers) {
	Repo = r
}

func (b *BHandlers) AddCSRFData(pd *models.PageData, r *http.Request) *models.PageData {
	pd.CSRFToken = nosurf.Token(r)

	if b.App.Session.Exists(r.Context(), "user_id") {
		pd.IsAuthenticated = 1
	}
	return pd
}

// HomeHandler - for getting the home page
func (b *BHandlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
	//ut := b.App.Session.Get(r.Context(), "user_type")
	//log.Println("user type: ", ut)

	pd := b.AddCSRFData(&models.PageData{}, r)

	var artList map[int]interface{}
	artList, err := b.DB.Get3BlogPost()
	if err != nil {
		log.Println(err)
		return
	}

	err = b.App.UITemplates.ExecuteTemplate(w, "home.page.tmpl", &models.PageData{
		Data:            artList,
		CSRFToken:       pd.CSRFToken,
		IsAuthenticated: pd.IsAuthenticated,
	})
	if err != nil {
		log.Println(err)
	}
}

// AboutHandler - for getting the about page
func (b *BHandlers) AboutHandler(w http.ResponseWriter, r *http.Request) {
	//strMap := make(map[string]string)
	//render.RenderTemplate(w, r, "about.page.tmpl", &models.PageData{StrMap: strMap})
	pd := b.AddCSRFData(&models.PageData{}, r)

	err := b.App.UITemplates.ExecuteTemplate(w, "about.page.tmpl", &models.PageData{
		CSRFToken:       pd.CSRFToken,
		IsAuthenticated: pd.IsAuthenticated})
	if err != nil {
		return
	}
}

// LoginHandler - for getting the login page
func (b *BHandlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	pd := b.AddCSRFData(&models.PageData{}, r)

	err := b.App.UITemplates.ExecuteTemplate(w, "login.page.tmpl", &models.PageData{
		CSRFToken:       pd.CSRFToken,
		IsAuthenticated: pd.IsAuthenticated})
	if err != nil {
		http.Error(w, "unable to execute the template", http.StatusInternalServerError)
		return
	}
}

// PageHandler - for getting the individual pages
func (b *BHandlers) PageHandler(w http.ResponseWriter, r *http.Request) {
	err := b.App.UITemplates.ExecuteTemplate(w, "page.page.tmpl", &models.PageData{})
	if err != nil {
		return
	}

}

// MakePostHandler - for creating new posts
func (b *BHandlers) MakePostHandler(w http.ResponseWriter, r *http.Request) {
	pd := b.AddCSRFData(&models.PageData{}, r)

	if !b.App.Session.Exists(r.Context(), "user_id") {
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
	err := b.App.UITemplates.ExecuteTemplate(w, "make-post.page.tmpl", &models.PageData{
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
func (b *BHandlers) PostMakePostHandler(w http.ResponseWriter, r *http.Request) {
	pd := b.AddCSRFData(&models.PageData{}, r)

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	uID := (b.App.Session.Get(r.Context(), "user_id")).(int)
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
		err := b.App.UITemplates.ExecuteTemplate(w, "make-post.page.tmpl", &models.PageData{
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
	err = b.DB.InsertPost(article)
	if err != nil {
		log.Fatal(err)
	}

	b.App.Session.Put(r.Context(), "article", article)
	http.Redirect(w, r, "/article-received", http.StatusSeeOther)
}

// ArticleReceived - get article
func (b *BHandlers) ArticleReceived(w http.ResponseWriter, r *http.Request) {
	pd := b.AddCSRFData(&models.PageData{}, r)

	article, ok := b.App.Session.Get(r.Context(), "article").(models.Article)
	if !ok {
		log.Println("Cant get data from session")
		b.App.Session.Put(r.Context(), "error", "Cant get data from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	data := make(map[string]interface{})
	data["article"] = article

	//render.RenderTemplate(w, r, "article-received.page.tmpl", &models.PageData{
	//	Data: data,
	//})
	err := b.App.UITemplates.ExecuteTemplate(w, "article-received.page.tmpl", &models.PageData{
		//Data:            data,
		CSRFToken:       pd.CSRFToken,
		IsAuthenticated: pd.IsAuthenticated,
	})
	if err != nil {
		return
	}
}

// PostLoginHandler - for getting the individual pages
func (b *BHandlers) PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	_ = b.App.Session.RenewToken(r.Context())
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
		err := b.App.UITemplates.ExecuteTemplate(w, "login.page.tmpl", &models.PageData{Form: form})
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	id, _, err := b.DB.AuthenticateUser(email, password)
	if err != nil {
		b.App.Session.Put(r.Context(), "error", "Invalid Email OR Password")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	b.App.Session.Put(r.Context(), "user_id", id)
	b.App.Session.Put(r.Context(), "flash", "Valid Login")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (b *BHandlers) PostAdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	_ = b.App.Session.RenewToken(r.Context())
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
		err := b.App.UITemplates.ExecuteTemplate(w, "login.page.tmpl", &models.PageData{Form: form})
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	id, _, err := b.DB.AuthenticateUser(email, password)
	if err != nil {
		b.App.Session.Put(r.Context(), "error", "Invalid Email OR Password")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		return
	}
	b.App.Session.Put(r.Context(), "user_id", id)
	b.App.Session.Put(r.Context(), "flash", "Valid Login")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (b *BHandlers) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	_ = b.App.Session.Destroy(r.Context())
	_ = b.App.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
