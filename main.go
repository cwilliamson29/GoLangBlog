package main

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/cwilliamson29/GoLangBlog/dbRepo"
	"github.com/cwilliamson29/GoLangBlog/handlers"
	"github.com/cwilliamson29/GoLangBlog/middleware"
	"github.com/cwilliamson29/GoLangBlog/models"
	"github.com/go-chi/chi/v5"
	mwc "github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
	"time"
)

var sessionManager *scs.SessionManager

//var app config.AppConfig

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

	dbc, err := dbRepo.ConnectSQL(dbRepo.DbConnection)
	if err != nil {
		log.Println("cant connect to dbRepo: ", err)
	}
	defer dbc.DB.Close()

	AdminTemplates := template.Must(template.ParseGlob("./templates/admin/*.tmpl"))
	UITemplates := template.Must(template.ParseGlob("./templates/ui/*.tmpl"))

	repo := handlers.NewRepo(dbc, AdminTemplates, UITemplates, sessionManager)
	handlers.NewHandlers(repo)

	router := chi.NewRouter()
	router.Use(mwc.Logger)
	router.Use(mwc.Recoverer)
	router.Use(middleware.LogRequestInfo)

	// Site Routes GET
	router.Get("/", handlers.Repo.HomeHandler)
	router.Get("/about", handlers.Repo.AboutHandler)
	router.Get("/login", handlers.Repo.LoginHandler)
	router.Get("/logout", handlers.Repo.LogoutHandler)
	router.Get("/makepost", handlers.Repo.MakePostHandler)
	router.Get("/article-received", handlers.Repo.ArticleReceived)
	router.Get("/page", handlers.Repo.PageHandler)

	// Site Routes POST
	router.Post("/makepost", handlers.Repo.PostMakePostHandler)
	router.Post("/login", handlers.Repo.PostLoginHandler)

	// ADMIN routes GET
	router.Get("/admin", handlers.Repo.AdminHandler)
	router.Get("/admin/login", handlers.Repo.AdminLoginHandler)
	router.Get("/admin/logout", handlers.Repo.LogoutHandler)
	router.Get("/admin/users", handlers.Repo.AdminUsersHandler)

	// ADMIN routes POST
	router.Post("/admin/login", handlers.Repo.PostAdminLoginHandler)
	router.Post("/admin/user/create", handlers.Repo.PostUserCreateHandler)
	router.Post("/admin/user/delete", handlers.Repo.PostUserDeleteHandler)

	fileServer := http.FileServer(http.Dir("./templates/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	port := "8080"

	log.Println("Listing on port: ", port)
	http.ListenAndServe(":"+port, sessionManager.LoadAndSave(router))

	//srv := &http.Server{
	//	Addr:    ":8080",
	//	Handler: router,
	//}
	//
	//err = srv.ListenAndServe()
	//if err != nil {
	//	log.Fatal(err)
	//}

}
