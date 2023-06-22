package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/forms"
	"github.com/cwilliamson29/GoLangBlog/models"
	"log"
	"net/http"
	"strconv"
)

func (b *BHandlers) PostUserCreateHandler(w http.ResponseWriter, r *http.Request) {
	pd := b.UserExists(&models.PageData{}, r)

	// Check if user logged in
	uAdmin, err := b.IsAdmin(w, r)
	if err != nil {
		log.Println(err)
	}
	// Check if user is admin
	if uAdmin {
		err = r.ParseForm()
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

		var userList map[int]interface{}
		userList, err := b.DB.GetAllUsers()
		if err != nil {
			log.Println(err)
			return
		}

		// Redirect back to users
		err2 := b.AdminTemplates.ExecuteTemplate(w, "admin.users.page.tmpl", &models.PageData{
			IsAuthenticated: pd.IsAuthenticated,
			Data:            userList,
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

// PostUserDeleteHandler - Deletes user from database
func (b *BHandlers) PostUserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	pd := b.UserExists(&models.PageData{}, r)

	// Check if user logged in
	uAdmin, err := b.IsAdmin(w, r)
	if err != nil {
		log.Println(err)
	}
	// Check if user is admin
	if uAdmin {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
			return
		}
		uId, _ := strconv.Atoi(r.Form.Get("user_id"))
		delType := r.Form.Get("del_type")
		uDel := make(map[string]any)
		// Write to the DB
		if delType == "delete" {
			err = b.DB.DeleteUser(uId)
			if err != nil {
				uDel["error"] = err
			} else {
				uDel["success"] = "User Deleted Successfully"
			}
		} else {
			uDel["success"] = "User Banned Successfully"
		}

		var userList map[int]interface{}
		userList, err = b.DB.GetAllUsers()
		if err != nil {
			log.Println(err)
			return
		}
		// Redirect back to users
		err2 := b.AdminTemplates.ExecuteTemplate(w, "admin.users.page.tmpl", &models.PageData{
			IsAuthenticated: pd.IsAuthenticated,
			Data:            userList,
			Active:          "users",
			UserDel:         uDel,
		})
		if err2 != nil {
			log.Println(err)
			return
		}
	}
}
