package handlers

import (
	"github.com/cwilliamson29/GoLangBlog/models"
	"log"
	"net/http"
)

func (b *BHandlers) IsAdmin(w http.ResponseWriter, r *http.Request) (bool, error) {
	pd := b.AddCSRFData(&models.PageData{}, r)
	// Check if user is admin
	if !b.App.Session.Exists(r.Context(), "user_id") {
		//render.RenderUnauthorizedTemplate(w, r, "authorizeLogin.page.tmpl", &models.PageData{})
		err := b.App.AdminTemplates.ExecuteTemplate(w, "admin.login.page.tmpl", &models.PageData{
			CSRFToken:       pd.CSRFToken,
			IsAuthenticated: pd.IsAuthenticated,
		})
		if err != nil {
			http.Error(w, "unable to execute the template", http.StatusInternalServerError)
			log.Println(err)
		}
	} else {
		uid := b.App.Session.Get(r.Context(), "user_id").(int)

		u, _ := b.DB.GetUserById(uid)

		//log.Println("user: ", u.UserType)
		if u.UserType != 3 {
			//log.Println("user_type: ", u.UserType, "and bool: ", u.UserType == 3)
			// Check if user is admin
			err := b.App.AdminTemplates.ExecuteTemplate(w, "unauthorized.page.tmpl", &models.PageData{
				CSRFToken:       pd.CSRFToken,
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
