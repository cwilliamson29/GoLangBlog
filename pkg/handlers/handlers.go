package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/models"
	"github.com/cwilliamson29/GoLangBlog/pkg/render"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.PageData{})
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {

	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = "This page is where we pass some data"

	render.RenderTemplate(w, "about.page.tmpl", &models.PageData{StrMap: strMap})

}
