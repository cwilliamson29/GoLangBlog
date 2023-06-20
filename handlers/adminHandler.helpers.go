package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/models"
	"log"
	"net/http"
)

func (b *BHandlers) IsAdmin(w http.ResponseWriter, r *http.Request) (bool, error) {
	pd := b.UserExists(&models.PageData{}, r)

	// Check if user is admin
	if !b.Session.Exists(r.Context(), "user_id") {
		err := b.AdminTemplates.ExecuteTemplate(w, "admin.login.page.tmpl", &models.PageData{
			IsAuthenticated: pd.IsAuthenticated,
		})
		if err != nil {
			http.Error(w, "unable to execute the template", http.StatusInternalServerError)
			log.Println(err)
		}
	} else {
		uid := b.Session.Get(r.Context(), "user_id").(int)
		u, _ := b.DB.GetUserById(uid)

		if u.UserType != 3 {
			// Check if user is admin
			err := b.AdminTemplates.ExecuteTemplate(w, "unauthorized.page.tmpl", &models.PageData{
				IsAuthenticated: pd.IsAuthenticated,
			})
			if err != nil {
				http.Error(w, "unable to execute the template", http.StatusInternalServerError)
				log.Println(err)
			}
		} else {
			return true, nil
		}
	}
	return false, nil
}
