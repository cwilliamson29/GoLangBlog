package routes

import (
	middleware2 "github.com/cwilliamson29/GoLangBlog/middleware"
	"github.com/cwilliamson29/GoLangBlog/pkg/config"
	"github.com/cwilliamson29/GoLangBlog/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Router(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware2.LogRequestInfo)

	mux.Use(middleware2.NoSurf)
	mux.Use(middleware2.SetupSession)

	// Site Routes GET
	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	mux.Get("/login", handlers.Repo.LoginHandler)
	mux.Get("/logout", handlers.Repo.LogoutHandler)
	mux.Get("/makepost", handlers.Repo.MakePostHandler)
	mux.Get("/article-received", handlers.Repo.ArticleReceived)
	mux.Get("/page", handlers.Repo.PageHandler)

	// Site Routes POST
	mux.Post("/makepost", handlers.Repo.PostMakePostHandler)
	mux.Post("/login", handlers.Repo.PostLoginHandler)

	// ADMIN routes GET
	mux.Get("/admin", handlers.Repo.AdminHandler)
	mux.Get("/admin/login", handlers.Repo.AdminLoginHandler)
	mux.Get("/admin/logout", handlers.Repo.LogoutHandler)
	mux.Get("/admin/users", handlers.Repo.AdminUsersHandler)

	// ADMIN routes POST
	mux.Post("/admin/login", handlers.Repo.PostAdminLoginHandler)
	mux.Post("/admin/user/create", handlers.Repo.PostUserCreateHandler)

	fileServer := http.FileServer(http.Dir("./templates/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	//mux.Handle("/templates/*", http.StripPrefix("./templates", fileServer))
	//mux.Handle("/css/*", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

	return mux
}
