package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/models"
	"github.com/cwilliamson29/GoLangBlog/pkg/render"
	"net/http"
)

// User type 1 - normal user
// User type 2 - moderator user
// User type 3 - admin user

// Unauthorized Handler
func (m *Repository) UnauthorizedHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	render.RenderTemplate(w, r, "unauthorized.page.tmpl", &models.PageData{StrMap: strMap})
}

// LoginHandler - for getting the login page
func (m *Repository) AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	render.RenderUnauthorizedTemplate(w, r, "authorizeLogin.page.tmpl", &models.PageData{StrMap: strMap})
}

// AdminHandler - for getting the admin page
func (m *Repository) AdminHandler(w http.ResponseWriter, r *http.Request) {
	uid := m.App.Session.Get(r.Context(), "user_id")
	id, _ := uid.(int)
	// Check if user logged in
	if !m.App.Session.Exists(r.Context(), "user_id") {
		render.RenderUnauthorizedTemplate(w, r, "authorizeLogin.page.tmpl", &models.PageData{})
	} else {
		u, _ := m.DB.GetUserById(id)
		//log.Println("user_type: ", u.UserType, "and bool: ", u.UserType == 3)
		// Check if user is admin
		if u.UserType != 3 {
			render.RenderUnauthorizedTemplate(w, r, "unauthorized.page.tmpl", &models.PageData{})
		} else {
			render.RenderAdminTemplate(w, r, "admin.page.tmpl", &models.PageData{})
		}
	}
}
