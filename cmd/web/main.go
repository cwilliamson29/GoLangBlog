package main

import (
	"database/sql"
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/cwilliamson29/GoLangBlog/models"
	"github.com/cwilliamson29/GoLangBlog/pkg/config"
	"github.com/cwilliamson29/GoLangBlog/pkg/dbdriver"
	"github.com/cwilliamson29/GoLangBlog/pkg/handlers"
	"html/template"
	"log"
	"net/http"
	"time"
)

var sessionManager *scs.SessionManager
var app config.AppConfig

func main() {
	//render.NewAppConfig(&app)
	gob.Register(models.Article{})
	gob.Register(models.User{})
	gob.Register(models.Post{})

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = false
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = sessionManager

	app.AdminTemplates = template.Must(template.ParseGlob("./templates/admin/*.tmpl"))
	app.UITemplates = template.Must(template.ParseGlob("./templates/ui/*.tmpl"))

	db, err := dbdriver.ConnectSQL("host=localhost port=5432 dbname=blog_db user=postgres password=TurtleDove")
	if err != nil {
		log.Fatal("cant connect to db: ", err)
	}

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	defer func(SQL *sql.DB) {
		err := SQL.Close()
		if err != nil {

		}
	}(db.SQL)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

//func run() (*dbdriver.DB, error) {
//
//}
