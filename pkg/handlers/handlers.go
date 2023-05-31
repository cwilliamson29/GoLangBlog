package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/models"
	"github.com/cwilliamson29/GoLangBlog/pkg/config"
	"github.com/cwilliamson29/GoLangBlog/pkg/render"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		App: ac,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// HomeHandler - for getting the home page
func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	m.App.Session.Put(r.Context(), "userid", "cwilliamson")
	render.RenderTemplate(w, r, "home.page.tmpl", &models.PageData{})
}

// AboutHandler - for getting the about page
func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "about.page.tmpl", &models.PageData{StrMap: strMap})
}

// LoginHandler - for getting the login page
func (m *Repository) LoginHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "login.page.tmpl", &models.PageData{StrMap: strMap})
}

// MakePostHandler - for creating new posts
func (m *Repository) MakePostHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "make-post.page.tmpl", &models.PageData{StrMap: strMap})
}

// PageHandler - for getting the individual pages
func (m *Repository) PageHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "page.page.tmpl", &models.PageData{StrMap: strMap})
}

// PostMakePostHandler - Post method for submitting new posts
func (m *Repository) PostMakePostHandler(w http.ResponseWriter, r *http.Request) {
	// strMap := make(map[string]string)
	blogTitle := r.Form.Get("blog_title")
	blogArticle := r.Form.Get("blog_article")
	w.Write([]byte(blogArticle))
	w.Write([]byte(blogTitle))
}
