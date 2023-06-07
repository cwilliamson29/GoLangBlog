package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/models"
	"github.com/cwilliamson29/GoLangBlog/pkg/render"
	"net/http"
)

// AdminHandler - for getting the admin page
func (m *Repository) AdminHandler(w http.ResponseWriter, r *http.Request) {
	if !m.App.Session.Exists(r.Context(), "user_id") {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	}

	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "admin.page.tmpl", &models.PageData{StrMap: strMap})
}
