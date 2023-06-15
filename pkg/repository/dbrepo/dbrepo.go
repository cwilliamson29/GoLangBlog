package dbrepo

import (
	"database/sql"
	"github.com/cwilliamson29/GoLangBlog/pkg/config"
	"github.com/cwilliamson29/GoLangBlog/pkg/repository"
)

type MySqlDB struct {
	App       *config.AppConfig
	DB        *sql.DB
	User      string
	Password  string
	Host      string
	Database  string
	ParseTime string
	err       error
}

func NewSQLRepo(conn *sql.DB, ac *config.AppConfig) repository.DatabaseRepo {
	return &MySqlDB{
		App:       ac,
		DB:        conn,
		User:      User,
		Password:  Password,
		Host:      Host,
		Database:  Database,
		ParseTime: ParseTime,
	}
}
