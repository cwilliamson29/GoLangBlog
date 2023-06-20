package dbRepo

import (
	"database/sql"
)

type MySqlDB struct {
	//App       *config.AppConfig
	DB        *sql.DB
	User      string
	Password  string
	Host      string
	Database  string
	ParseTime string
	err       error
}

// , ac *config.AppConfig
func NewSQLRepo(conn *sql.DB) DatabaseRepo {
	return &MySqlDB{
		//App:       ac,
		DB:        conn,
		User:      User,
		Password:  Password,
		Host:      Host,
		Database:  Database,
		ParseTime: ParseTime,
	}
}
