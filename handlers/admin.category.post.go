package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/models"
	"log"
	"net/http"
	"strconv"
)

func (b *BHandlers) PostCategoryAddHandler(w http.ResponseWriter, r *http.Request) {
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
		n := r.Form.Get("name")
		stat := make(map[string]any)
		// Write to the DB
		err = b.DB.CateAdd(n)
		if err != nil {
			stat["catAddError"] = err
		} else {
			stat["catAddSuccess"] = "Category Added Successfully"
		}
		b.CategoryTempExecute(w, stat, "addc")
	}
}

func (b *BHandlers) PostSubCategoryAddHandler(w http.ResponseWriter, r *http.Request) {
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
		id, _ := strconv.Atoi(r.Form.Get("category_id"))
		n := r.Form.Get("name")
		stat := make(map[string]any)
		// Write to the DB
		err = b.DB.SubCateAdd(n, id)
		if err != nil {
			stat["scatAddError"] = err
		} else {
			stat["scatAddSuccess"] = "Category Added Successfully"
		}
		b.CategoryTempExecute(w, stat, "addsc")
	}
}

// PostCategoryDeleteHandler - Delete category from database
func (b *BHandlers) PostCategoryDeleteHandler(w http.ResponseWriter, r *http.Request) {
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
		id, _ := strconv.Atoi(r.Form.Get("category_id"))
		sr := r.Form.Get("subRemove")
		stat := make(map[string]any)
		// Write to the DB
		if sr == "true" {
			err = b.DB.SubCateAdd(n, id)
			if err != nil {
				stat["catDelError"] = err
			} else {
				stat["catDelSuccess"] = "Category Removed Successfully"
			}
		} else {
			stat["catDelError"] = ""
		}

		b.CategoryTempExecute(w, stat, "addsc")
	}
}
func (b *BHandlers) CategoryTempExecute(w http.ResponseWriter, cAdd map[string]any, ca string) {
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
	// Redirect back to users
	err2 := b.AdminTemplates.ExecuteTemplate(w, "admin.category.page.tmpl", &models.PageData{
		//IsAuthenticated: pd.IsAuthenticated,
		Data:   cList,
		Data2:  scList,
		Active: "category",
		Status: cAdd,
		CA:     ca,
	})
	if err2 != nil {
		log.Println(err)
		return
	}
}
