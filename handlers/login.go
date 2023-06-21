package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/forms"
	"github.com/cwilliamson29/GoLangBlog/models"
	"log"
	"net/http"
)

// LoginHandler - for getting the login page
func (b *BHandlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	pd := b.UserExists(&models.PageData{}, r)

	err := b.UITemplates.ExecuteTemplate(w, "login.page.tmpl", &models.PageData{
		IsAuthenticated: pd.IsAuthenticated,
	})
	if err != nil {
		http.Error(w, "unable to execute the template", http.StatusInternalServerError)
		return
	}
}

// AdminLoginHandler - for getting the login page
func (b *BHandlers) AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	err := b.UITemplates.ExecuteTemplate(w, "authorizeLogin.page.tmpl", &models.PageData{})
	if err != nil {
		return
	}
}

// PostLoginHandler - for getting the individual pages
func (b *BHandlers) PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	err := b.LoginHelper("ui", w, r)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// LoginHelper - Logic for login function
func (b *BHandlers) LoginHelper(t string, w http.ResponseWriter, r *http.Request) error {
	_ = b.Session.RenewToken(r.Context())
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	form := forms.New(r.PostForm)
	form.HasRequired("email", "password")
	form.IsEmail("email")

	if !form.Valid() {
		err := b.UITemplates.ExecuteTemplate(w, "login.page.tmpl", &models.PageData{Form: form})
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
	id, _, err := b.DB.AuthenticateUser(email, password)
	if err != nil {
		b.Session.Put(r.Context(), "error", "Invalid Email OR Password")
		if t == "ui" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else if t == "admin" {
			http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
		}
		return nil
	}

	b.Session.Put(r.Context(), "user_id", id)
	b.Session.Put(r.Context(), "flash", "Valid Login")

	return nil
}

func (b *BHandlers) PostAdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	err := b.LoginHelper("admin", w, r)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

// LogoutHandler - Handle logout
func (b *BHandlers) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	_ = b.Session.Destroy(r.Context())
	_ = b.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
