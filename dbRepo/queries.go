package dbRepo

import "errors"

// Post Queries
var queryInsertPost = `INSERT INTO posts(title, content, user_id) VALUES($1, $2, $3)`
var queryGet3BlogPosts = `SELECT id, user_id, title, content FROM posts ORDER BY id DESC LIMIT ?`

// User Queries
var queryGetAllUsers = `SELECT id, name, email, user_type FROM users ORDER BY id DESC`
var queryGetUserById = `SELECT name, email, password, acct_created, last_login, user_type, id FROM users WHERE id = ?`
var queryFindByEmail = `SELECT email FROM users WHERE email = ?`
var queryAddUser = `INSERT INTO users(name, email, password, user_type, acct_created, last_login) VALUES(?, ?, ?, ?, ?, ?)`
var queryDeleteUser = `DELETE FROM users WHERE id = ?`

func (m *MySqlDB) DeleteUser(id int) error {
	var success bool
	ct := m.Connect()
	if ct {
		success = m.Delete(queryDeleteUser, id)
	}
	// check if return true for success
	if success {
		return nil
	} else {
		return errors.New("user not deleted")
	}
}
