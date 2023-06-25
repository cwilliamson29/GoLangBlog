package dbRepo

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

// Get - Get method for assisting in dbRepo function
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

// Delete Delete
func (m *MySqlDB) Delete(query string, args ...any) bool {
	var success = false
	var stmt *sql.Stmt
	stmt, m.err = m.DB.Prepare(query)
	if m.err != nil {
		log.Println("Error:", m.err.Error())
	} else {
		defer stmt.Close()
		res, err := stmt.Exec(args...)
		if err != nil {
			log.Println("Delete Exec err:", err.Error())
		} else {
			affectedRows, _ := res.RowsAffected()
			if affectedRows == 0 {
				log.Println("Error: No records deleted")
			} else {
				success = true
			}
		}
	}
	return success
}

// Update Update
func (m *MySqlDB) Update(query string, args ...any) bool {
	var success = false
	var stmtUp *sql.Stmt
	stmtUp, m.err = m.DB.Prepare(query)
	if m.err != nil {
		log.Println("Error:", m.err.Error())
	} else {
		defer stmtUp.Close()
		res, err := stmtUp.Exec(args...)
		if err != nil {
			log.Println("Update Exec err:", err.Error())
		} else {
			affectedRows, _ := res.RowsAffected()
			if affectedRows == 0 {
				log.Println("Error: No records updated")
			} else {
				success = true
			}
		}
	}
	return success
}

// Insert Insert
func (m *MySqlDB) Insert(query string, args ...any) (bool, int64) {
	var success = false
	var id int64 = -1
	var stmtIns *sql.Stmt
	stmtIns, m.err = m.DB.Prepare(query)
	if m.err != nil {
		log.Println("Error:", m.err.Error())
	} else {
		defer stmtIns.Close()
		res, err := stmtIns.Exec(args...)
		if err != nil {
			log.Println("Insert Exec err:", err.Error())
		} else {
			id, err = res.LastInsertId()
			affectedRows, _ := res.RowsAffected()
			if err == nil && affectedRows > 0 {
				success = true
			}
		}
	}
	return success, id
}
