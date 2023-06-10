package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/models"
	"log"
	"net/http"
)

// User type 1 - normal user
// User type 2 - moderator user
// User type 3 - admin user

// Unauthorized Handler
func (m *Repository) UnauthorizedHandler(w http.ResponseWriter, r *http.Request) {
	pd := m.AddCSRFData(&models.PageData{}, r)

	//strMap := make(map[string]string)
	//render.RenderTemplate(w, r, "unauthorized.page.tmpl", &models.PageData{StrMap: strMap})
	err := m.App.AdminTemplates.ExecuteTemplate(w, "admin.home.page.tmpl", &models.PageData{
		CSRFToken:       pd.CSRFToken,
		IsAuthenticated: pd.IsAuthenticated,
	})
	if err != nil {
		return
	}
}

// LoginHandler - for getting the login page
func (m *Repository) AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	//strMap := make(map[string]string)
	//render.RenderUnauthorizedTemplate(w, r, "authorizeLogin.page.tmpl", &models.PageData{StrMap: strMap})
	err := m.App.UITemplates.ExecuteTemplate(w, "authorizeLogin.page.tmpl", &models.PageData{})
	if err != nil {
		return
	}
}

// AdminHandler - for getting the admin page
func (m *Repository) AdminHandler(w http.ResponseWriter, r *http.Request) {
	pd := m.AddCSRFData(&models.PageData{}, r)

	uid := m.App.Session.Get(r.Context(), "user_id")
	id, _ := uid.(int)
	// Check if user logged in
	if !m.App.Session.Exists(r.Context(), "user_id") {
		//render.RenderUnauthorizedTemplate(w, r, "authorizeLogin.page.tmpl", &models.PageData{})
		err := m.App.UITemplates.ExecuteTemplate(w, "login.page.tmpl", &models.PageData{
			CSRFToken:       pd.CSRFToken,
			IsAuthenticated: pd.IsAuthenticated,
		})
		if err != nil {
			http.Error(w, "unable to execute the template", http.StatusInternalServerError)
			return
		}
	} else {
		u, _ := m.DB.GetUserById(id)
		//log.Println("user_type: ", u.UserType, "and bool: ", u.UserType == 3)
		// Check if user is admin
		if u.UserType != 2 {
			//render.RenderUnauthorizedTemplate(w, r, "unauthorized.page.tmpl", &models.PageData{})
			err := m.App.AdminTemplates.ExecuteTemplate(w, "unauthorized.page.tmpl", &models.PageData{
				CSRFToken:       pd.CSRFToken,
				IsAuthenticated: pd.IsAuthenticated,
			})
			if err != nil {
				return
			}
		} else {
			log.Println("here")
			//render.RenderAdminTemplate(w, r, "admin.page.tmpl", &models.PageData{})
			err := m.App.AdminTemplates.ExecuteTemplate(w, "admin.home.page.tmpl", &models.PageData{
				CSRFToken:       pd.CSRFToken,
				IsAuthenticated: pd.IsAuthenticated,
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
