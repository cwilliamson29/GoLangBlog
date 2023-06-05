package helpers

import (
	"github.com/cwilliamson29/GoLangBlog/pkg/config"
	"log"
	"net/http"
)

var app config.AppConfig

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// IsAuthenticated - Checks if user exists
func IsAuthenticated(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "user_id")
	return exists
}
