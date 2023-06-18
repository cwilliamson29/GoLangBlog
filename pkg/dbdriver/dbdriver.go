package dbdriver

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConns = 20
const maxIdleDbConns = 10
const maxDbLifetime = 5 * time.Minute

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}

func ConnectSQL(dsn string) (*DB, error) {
	db, err := NewDatabase(dsn)
	if err != nil {
		log.Println(err)
	}
	db.SetMaxOpenConns(maxOpenDbConns)
	db.SetConnMaxIdleTime(maxIdleDbConns)
	db.SetConnMaxIdleTime(maxDbLifetime)

	dbConn.SQL = db
	return dbConn, nil
}
