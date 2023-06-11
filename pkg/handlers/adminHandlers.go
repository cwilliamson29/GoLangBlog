package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/models"
	"log"
	"net/http"
)

// User type 1 - normal user
// User type 2 - moderator user
// User type 3 - admin user

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

	// Check if user logged in
	uAdmin, err := m.IsAdmin(w, r)
	if err != nil {
		log.Println(err)
	}

	if uAdmin == true {
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
