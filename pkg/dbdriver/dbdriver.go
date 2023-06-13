package dbdriver

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/jackc/pgx/pgconn
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
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
		return nil, err
	}

	//db, err := pgx.Connect(context.Background(), dsn)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectSQL(dsn string) (*DB, error) {
	db, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(maxOpenDbConns)
	db.SetConnMaxIdleTime(maxIdleDbConns)
	db.SetConnMaxIdleTime(maxDbLifetime)

	dbConn.SQL = db
	return dbConn, nil
}
