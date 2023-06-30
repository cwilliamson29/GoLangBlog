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

func (b *BHandlers) PostMenuEditIsNavHandler(w http.ResponseWriter, r *http.Request) {
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
		nav, _ := strconv.Atoi(r.Form.Get("is_navbar"))
		id, _ := strconv.Atoi(r.Form.Get("menu_id"))

		Stat := make(map[string]interface{})

		// Write to the DB
		suc := b.DB.UpdateIsNav(nav, id)
		if !suc {
			Stat["isNavError"] = err
		} else {
			Stat["isNavSuccess"] = "Main Navbar Changed Successfully"
		}
		b.MenuTempExecute(w, Stat, "menuCreate")
	}
}
func (b *BHandlers) PostMenuDeleteHandler(w http.ResponseWriter, r *http.Request) {
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
		id, _ := strconv.Atoi(r.Form.Get("menu_id"))

		Stat := make(map[string]interface{})

		// Write to the DB
		err := b.DB.DeleteMenuById(id)
		if err != nil {
			Stat["delNavError"] = err
		} else {
			Stat["delNavSuccess"] = "Navbar Deleted Successfully"
		}
		b.MenuTempExecute(w, Stat, "menuCreate")
	}
}

func (b *BHandlers) PostMenuCreateLinkHandler(w http.ResponseWriter, r *http.Request) {
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
		pId, _ := strconv.Atoi(r.Form.Get("parent_menu"))
		name := r.Form.Get("link_name")
		target := r.Form.Get("target")
		pos, _ := strconv.Atoi(r.Form.Get("position"))

		Stat := make(map[string]interface{})

		// Write to the DB
		err = b.DB.MenuLinkCreate(pId, name, target, pos)
		if err != nil {
			Stat["createLinkError"] = err
		} else {
			Stat["createLinkSuccess"] = "Link Added Successfully"
		}
		b.MenuTempExecute(w, Stat, "menuAddLink")
	}
}

func (b *BHandlers) PostMenuLinkDeleteHandler(w http.ResponseWriter, r *http.Request) {
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
		id, _ := strconv.Atoi(r.Form.Get("menu_id"))

		Stat := make(map[string]interface{})

		// Write to the DB
		err := b.DB.DeleteMenuById(id)
		if err != nil {
			Stat["delNavError"] = err
		} else {
			Stat["delNavSuccess"] = "Navbar Deleted Successfully"
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
	//log.Println("****status***", Stat)
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
