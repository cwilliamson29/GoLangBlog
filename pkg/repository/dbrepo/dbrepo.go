package dbrepo

import (
	"database/sql"
	"github.com/cwilliamson29/GoLangBlog/pkg/config"
	"github.com/cwilliamson29/GoLangBlog/pkg/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostGresRepo(conn *sql.DB, ac *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: ac,
		DB:  conn,
	}
}
