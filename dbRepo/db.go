package dbRepo

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
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

//var dbConn = &MySqlDB{}

const maxOpenDbConns = 20
const maxIdleDbConns = 10
const maxDbLifetime = 5 * time.Minute

// , ac *config.AppConfig
func NewSQLRepo(conn *sql.DB) DatabaseRepo {
	return &MySqlDB{
		DB:        conn,
		User:      User,
		Password:  Password,
		Host:      Host,
		Database:  Database,
		ParseTime: ParseTime,
	}
}

func ConnectSQL(dsn string) (*MySqlDB, error) {
	//db, err := NewDatabase(dsn)
	var dbConn = &MySqlDB{}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenDbConns)
	db.SetConnMaxIdleTime(maxIdleDbConns)
	db.SetConnMaxIdleTime(maxDbLifetime)

	dbConn.DB = db
	return dbConn, nil
}
