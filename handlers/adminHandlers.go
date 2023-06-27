package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/models"
	"log"
	"net/http"
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
		var userList map[int]interface{}
		userList, err := b.DB.GetAllUsers()
		if err != nil {
			log.Println(err)
			return
		}

		err = b.AdminTemplates.ExecuteTemplate(w, "admin.users.page.tmpl", &models.PageData{
			IsAuthenticated: pd.IsAuthenticated,
			Data:            userList,
			Active:          "users",
			UA:              "userAdd",
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
}
func (b *BHandlers) AdminMenuHandler(w http.ResponseWriter, r *http.Request) {
	pd := b.UserExists(&models.PageData{}, r)

	// Check if user logged in
	uAdmin, err := b.IsAdmin(w, r)
	if err != nil {
		log.Println(err)
	}

	// Check if user is admin
	if uAdmin {

		err = b.AdminTemplates.ExecuteTemplate(w, "admin.menu.page.tmpl", &models.PageData{
			IsAuthenticated: pd.IsAuthenticated,
			Active:          "menu",
			MA:              "menuCreate",
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
}
func (b *BHandlers) AdminCategoryHandler(w http.ResponseWriter, r *http.Request) {
	pd := b.UserExists(&models.PageData{}, r)

	// Check if user logged in
	uAdmin, err := b.IsAdmin(w, r)
	if err != nil {
		log.Println(err)
	}

	// Check if user is admin
	if uAdmin {
		var cList map[int]interface{}
		var scList map[int]interface{}
		cList, err := b.DB.GetAllCategories()
		if err != nil {
			log.Println(err)
			return
		}
		scList, err = b.DB.GetAllSubCategories()
		if err != nil {
			log.Println(err)
			return
		}
		err = b.AdminTemplates.ExecuteTemplate(w, "admin.category.page.tmpl", &models.PageData{
			IsAuthenticated: pd.IsAuthenticated,
			Data:            cList,
			Data2:           scList,
			Active:          "categories",
			CA:              "addc",
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
}
