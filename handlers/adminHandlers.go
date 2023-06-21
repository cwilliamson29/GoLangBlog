package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/forms"
	"github.com/cwilliamson29/GoLangBlog/models"
	"log"
	"net/http"
	"strconv"
)

// User type 1 - normal user
// User type 2 - moderator user
// User type 3 - admin user

// AdminHandler - for getting the admin page
func (b *BHandlers) AdminHandler(w http.ResponseWriter, r *http.Request) {
	pd := b.UserExists(&models.PageData{}, r)

	// Check if user logged in
	uAdmin, err := b.IsAdmin(w, r)
	if err != nil {
		log.Println(err)
	}
	// Check if user is admin
	if uAdmin {
		err := b.AdminTemplates.ExecuteTemplate(w, "admin.home.page.tmpl", &models.PageData{
			IsAuthenticated: pd.IsAuthenticated,
			Active:          "home",
		})
		if err != nil {
			log.Println(err)
			return
		}
	}

}

func (b *BHandlers) AdminUsersHandler(w http.ResponseWriter, r *http.Request) {
	pd := b.UserExists(&models.PageData{}, r)

	// Check if user logged in
	uAdmin, err := b.IsAdmin(w, r)
	if err != nil {
		log.Println(err)
	}

	// Check if user is admin
	if uAdmin {
		err := b.AdminTemplates.ExecuteTemplate(w, "admin.users.page.tmpl", &models.PageData{
			IsAuthenticated: pd.IsAuthenticated,
			Active:          "users",
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (b *BHandlers) PostUserCreateHandler(w http.ResponseWriter, r *http.Request) {
	pd := b.UserExists(&models.PageData{}, r)

	// Check if user logged in
	uAdmin, err := b.IsAdmin(w, r)
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

		form := forms.New(r.PostForm)

		form.HasRequired("name", "email", "password")
		form.MinLength("name", 5, r)
		form.MinLength("password", 5, r)
		form.IsEmail("password")
		userAdd := make(map[string]interface{})

		// Write to the DB
		err = b.DB.AddUser(createUser)
		if err != nil {
			userAdd["error"] = err
		} else {
			userAdd["success"] = "User Added Successfully"
		}

		// Redirect back to users
		err2 := b.AdminTemplates.ExecuteTemplate(w, "admin.users.page.tmpl", &models.PageData{
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
