package helpers

import (
	"github.com/cwilliamson29/GoLangBlog/handlers"
	"log"
	"net/http"
)

var app *handlers.BHandlers

// User type 1 - normal user
// User type 2 - moderator user
// User type 3 - admin user

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
