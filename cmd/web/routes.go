package main

import (
	"github.com/cwilliamson29/GoLangBlog/pkg/config"
	"github.com/cwilliamson29/GoLangBlog/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(LogRequestInfo)

	mux.Use(NoSurf)
	mux.Use(SetupSession)

	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	mux.Get("/login", handlers.Repo.LoginHandler)
	mux.Get("/logout", handlers.Repo.LogoutHandler)
	mux.Post("/login", handlers.Repo.PostLoginHandler)
	mux.Get("/makepost", handlers.Repo.MakePostHandler)
	mux.Post("/makepost", handlers.Repo.PostMakePostHandler)
	mux.Get("/article-received", handlers.Repo.ArticleReceived)
	mux.Get("/page", handlers.Repo.PageHandler)

	// ADMIN routes
	mux.Get("/admin", handlers.Repo.AdminHandler)
	mux.Get("/admin/login", handlers.Repo.AdminLoginHandler)
	mux.Post("/admin/login", handlers.Repo.PostAdminLoginHandler)
	mux.Get("/admin/logout", handlers.Repo.LogoutHandler)

	fileServer := http.FileServer(http.Dir("./templates/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	//mux.Handle("/templates/*", http.StripPrefix("./templates", fileServer))
	//mux.Handle("/css/*", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

	return mux
}
