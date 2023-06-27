package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/models"
	"log"
	"net/http"
	"strconv"
)

func (b *BHandlers) PostMenuCreateHandler(w http.ResponseWriter, r *http.Request) {
	//pd := b.UserExists(&models.PageData{}, r)

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
		name := r.Form.Get("menu_name")
		nav, _ := strconv.Atoi(r.Form.Get("is_navbar"))

		Stat := make(map[string]interface{})

		// Write to the DB
		err = b.DB.MenuCreate(name, nav)
		if err != nil {
			Stat["createNavError"] = err
		} else {
			Stat["createNavSuccess"] = "Menu Added Successfully"
		}
		b.MenuTempExecute(w, Stat, "menuCreate")
	}
}

func (b *BHandlers) MenuTempExecute(w http.ResponseWriter, Stat map[string]any, ma string) {
	var menuList map[int]interface{}
	menuList, err := b.DB.GetAllMenus()
	if err != nil {
		log.Println(err)
		return
	}
	// Redirect back to menu
	err2 := b.AdminTemplates.ExecuteTemplate(w, "admin.menu.page.tmpl", &models.PageData{
		//IsAuthenticated: pd.IsAuthenticated,
		Data:   menuList,
		Active: "menu",
		Status: Stat,
		MA:     ma,
	})
	if err2 != nil {
		log.Println(err)
		return
	}
}
