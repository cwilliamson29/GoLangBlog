package handlers

import (
	"github.com/alexedwards/scs/v2"
	dbRepo2 "github.com/cwilliamson29/GoLangBlog/dbRepo"
	"github.com/cwilliamson29/GoLangBlog/forms"
	"github.com/cwilliamson29/GoLangBlog/models"
	"html/template"
	"log"
	"net/http"
)

type BHandlers struct {
	User           models.User
	DB             dbRepo2.DatabaseRepo
	InfoLog        *log.Logger
	Session        *scs.SessionManager
	CSRFToken      string
	AdminTemplates *template.Template
	UITemplates    *template.Template
}

var Repo *BHandlers

// ac *config.AppConfig,
func NewRepo(dbc *dbRepo2.MySqlDB, at *template.Template, ui *template.Template, ses *scs.SessionManager) *BHandlers {
	return &BHandlers{
		DB:             dbRepo2.NewSQLRepo(dbc.DB),
		Session:        ses,
		AdminTemplates: at,
		UITemplates:    ui,
	}
}

func NewHandlers(r *BHandlers) {
	Repo = r
}

func (b *BHandlers) UserExists(pd *models.PageData, r *http.Request) *models.PageData {
	if b.Session.Exists(r.Context(), "user_id") {
		pd.IsAuthenticated = 1
	}
	return pd
}

// HomeHandler - for getting the home page
func (b *BHandlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
	pd := b.UserExists(&models.PageData{}, r)

	var artList map[int]interface{}
	artList, err := b.DB.Get3BlogPost()
	if err != nil {
		log.Println(err)
		return
	}

	err = b.UITemplates.ExecuteTemplate(w, "home.page.tmpl", &models.PageData{
		Data:            artList,
		IsAuthenticated: pd.IsAuthenticated,
	})
	if err != nil {
		log.Println(err)
	}
}

// AboutHandler - for getting the about page
func (b *BHandlers) AboutHandler(w http.ResponseWriter, r *http.Request) {
	pd := b.UserExists(&models.PageData{}, r)

	err := b.UITemplates.ExecuteTemplate(w, "about.page.tmpl", &models.PageData{
		IsAuthenticated: pd.IsAuthenticated,
	})
	if err != nil {
		return
	}
}

// PageHandler - for getting the individual pages
func (b *BHandlers) PageHandler(w http.ResponseWriter, r *http.Request) {
	err := b.UITemplates.ExecuteTemplate(w, "page.page.tmpl", &models.PageData{})
	if err != nil {
		return
	}

}

// MakePostHandler - for creating new posts
func (b *BHandlers) MakePostHandler(w http.ResponseWriter, r *http.Request) {
	//pd := b.AddCSRFData(&models.PageData{}, r)

	//if !b.Session.Exists(r.Context(), "user_id") {
	//	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	//}
	var emptyArticle models.Article
	data := make(map[string]interface{})
	data["article"] = emptyArticle

	//render.RenderTemplate(w, r, "make-post.page.tmpl", &models.PageData{
	//	Form:            forms.New(nil),
	//	Data:            data,
	//	CSRFToken:       pd.CSRFToken,
	//	IsAuthenticated: pd.IsAuthenticated,
	//})
	err := b.UITemplates.ExecuteTemplate(w, "make-post.page.tmpl", &models.PageData{
		Form: forms.New(nil),
		//Data:            data,
		//CSRFToken:       pd.CSRFToken,
		//IsAuthenticated: pd.IsAuthenticated,
	})
	if err != nil {
		return
	}

}

// PostMakePostHandler - Post method for submitting new posts
func (b *BHandlers) PostMakePostHandler(w http.ResponseWriter, r *http.Request) {
	//pd := b.AddCSRFData(&models.PageData{}, r)

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	//uID := (b.Session.Get(r.Context(), "user_id")).(int)
	article := models.Post{
		Title:   r.Form.Get("blog_title"),
		Content: r.Form.Get("blog_article"),
		//UserID:  int(uID),
	}

	form := forms.New(r.PostForm)

	form.HasRequired("blog_title", "blog_article")

	form.MinLength("blog_title", 5, r)
	form.MinLength("blog_article", 5, r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["article"] = article
		err := b.UITemplates.ExecuteTemplate(w, "make-post.page.tmpl", &models.PageData{
			Form: form,
			//Data:            data,
			//CSRFToken:       pd.CSRFToken,
			//IsAuthenticated: pd.IsAuthenticated,
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

	//b.Session.Put(r.Context(), "article", article)
	http.Redirect(w, r, "/article-received", http.StatusSeeOther)
}

// ArticleReceived - get article
func (b *BHandlers) ArticleReceived(w http.ResponseWriter, r *http.Request) {
	//pd := b.AddCSRFData(&models.PageData{}, r)

	//article, ok := b.App.Session.Get(r.Context(), "article").(models.Article)
	//if !ok {
	//	log.Println("Cant get data from session")
	//	b.App.Session.Put(r.Context(), "error", "Cant get data from session")
	//	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	//	return
	//}
	//data := make(map[string]interface{})
	//data["article"] = article

	//render.RenderTemplate(w, r, "article-received.page.tmpl", &models.PageData{
	//	Data: data,
	//})
	err := b.UITemplates.ExecuteTemplate(w, "article-received.page.tmpl", &models.PageData{
		//Data:            data,
		//CSRFToken:       pd.CSRFToken,
		//IsAuthenticated: pd.IsAuthenticated,
	})
	if err != nil {
		return
	}
}
