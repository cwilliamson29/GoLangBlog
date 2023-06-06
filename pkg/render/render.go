package render

import (
	"fmt"
	"github.com/cwilliamson29/GoLangBlog/models"
	"github.com/cwilliamson29/GoLangBlog/pkg/config"
	"github.com/justinas/nosurf"
	"html/template"
	"net/http"
)

var tmplCache = make(map[string]*template.Template)

var app *config.AppConfig

func NewAppConfig(a *config.AppConfig) {
	app = a
}

func AddCSRFData(pd *models.PageData, r *http.Request) *models.PageData {
	pd.CSRFToken = nosurf.Token(r)

	if app.Session.Exists(r.Context(), "user_id") {
		pd.IsAuthenticated = 1
	}
	return pd
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, t string, pd *models.PageData) {
	var tmpl *template.Template
	var err error
	_, inMap := tmplCache[t]
	if !inMap {
		err = makeTemplateCache(t)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Template in cache")
		}
	}
	tmpl = tmplCache[t]

	pd = AddCSRFData(pd, r)

	err = tmpl.Execute(w, pd)
	if err != nil {
		fmt.Println(err)
	}
}

func makeTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tmplCache[t] = tmpl
	return nil
}
