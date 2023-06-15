package dbrepo

import (
	"database/sql"
	"log"
)

// DbRow array of database rows
type DbRow struct {
	Column []string
	Row    []string
}

// Connect - Test connection to the database
func (m *MySqlDB) Connect() bool {
	var rtn = false
	var conStr = m.User + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Database + m.ParseTime
	m.DB, m.err = sql.Open("mysql", conStr)
	if m.err == nil {
		m.err = m.DB.Ping()
		if m.err != nil {
			log.Println("Database Connect Error:", m.err.Error())
		} else {
			rtn = true
		}
	}
	return rtn
}

// Close - Close the database
func (m *MySqlDB) Close() bool {
	var rtn = false
	err := m.DB.Close()
	if err == nil {
		rtn = true
	}
	return rtn
}
