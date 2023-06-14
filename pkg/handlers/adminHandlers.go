package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/models"
	"github.com/cwilliamson29/GoLangBlog/pkg/forms"
	"log"
	"net/http"
	"strconv"
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
	// Check if user is admin
	if uAdmin {
		err := m.App.AdminTemplates.ExecuteTemplate(w, "admin.home.page.tmpl", &models.PageData{
			CSRFToken:       pd.CSRFToken,
			IsAuthenticated: pd.IsAuthenticated,
			Active:          "home",
		})
		if err != nil {
			log.Println(err)
			return
		}
	}

}

func (m *Repository) AdminUsersHandler(w http.ResponseWriter, r *http.Request) {
	pd := m.AddCSRFData(&models.PageData{}, r)

	// Check if user logged in
	uAdmin, err := m.IsAdmin(w, r)
	if err != nil {
		log.Println(err)
	}
	// Check if user is admin
	if uAdmin {
		err := m.App.AdminTemplates.ExecuteTemplate(w, "admin.users.page.tmpl", &models.PageData{
			CSRFToken:       pd.CSRFToken,
			IsAuthenticated: pd.IsAuthenticated,
			Active:          "users",
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (m *Repository) PostUserCreateHandler(w http.ResponseWriter, r *http.Request) {
	pd := m.AddCSRFData(&models.PageData{}, r)
	// Check if user logged in
	uAdmin, err := m.IsAdmin(w, r)
	if err != nil {
		log.Println(err)
	}
	// Check if user is admin
	if uAdmin {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			return
		}
		ut, _ := strconv.Atoi(r.Form.Get("user_type"))
		createUser := models.User{
			Name:     r.Form.Get("name"),
			Email:    r.Form.Get("email"),
			Password: r.Form.Get("password"),
			UserType: ut,
		}
		//log.Printf("name: %d, Email: %d, password: %d, userType: %d \n", createUser.Name, createUser.Email, createUser.Password, createUser.UserType)

		form := forms.New(r.PostForm)

		form.HasRequired("name", "email", "password")
		form.MinLength("name", 5, r)
		form.MinLength("password", 5, r)
		form.IsEmail("password")
		userAdd := make(map[string]interface{})

		// Write to the DB
		err = m.DB.AddUser(createUser)
		if err != nil {
			userAdd["error"] = err
		} else {
			userAdd["success"] = "User Added Successfully"
		}

		// Redirect back to users
		err2 := m.App.AdminTemplates.ExecuteTemplate(w, "admin.users.page.tmpl", &models.PageData{
			CSRFToken:       pd.CSRFToken,
			IsAuthenticated: pd.IsAuthenticated,
			Active:          "users",
			UserAdd:         userAdd,
		})
		if err2 != nil {
			log.Println(err)
			return
		}
	} else {
		log.Println("entered else")
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}
}
