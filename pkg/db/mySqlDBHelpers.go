package db

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

// Get - Get method for assisting in db function
func (m *MySqlDB) Get(query string, args ...any) *DbRow {
	var results DbRow
	qGet, err := m.DB.Prepare(query)
	if err != nil {
		log.Println(err)
	} else {
		defer qGet.Close()
		rows, err := qGet.Query(args...)
		if err != nil {
			log.Println(err)
		} else {
			defer rows.Close()
			col, err := rows.Columns()
			if err == nil {
				results.Column = col
				rowValues := make([]sql.RawBytes, len(col))
				scanArgs := make([]any, len(rowValues))
				for i := range rowValues {
					scanArgs[i] = &rowValues[i]
				}
				for rows.Next() {
					rows.Scan(scanArgs...)
					for _, cols := range rowValues {
						var value string
						if cols == nil {
							value = "NULL"
						} else {
							value = string(cols)
						}
						results.Row = append(results.Row, value)
					}
				}
			}
		}

	}
	return &results
}
