package dbRepo

// Post Queries
var queryInsertPost = `INSERT INTO posts(title, content, user_id) VALUES($1, $2, $3)`
var queryGet3BlogPosts = `SELECT id, user_id, title, content FROM posts ORDER BY id DESC LIMIT ?`

// User Queries
var queryLoginUser = `SELECT id, password FROM users WHERE email=?`
var queryGetAllUsers = `SELECT id, name, email, user_type, banned FROM users ORDER BY id DESC`
var queryGetUserById = `SELECT name, email, password, acct_created, last_login, user_type, id FROM users WHERE id = ?`
var queryFindByEmail = `SELECT email FROM users WHERE email = ?`
var queryAddUser = `INSERT INTO users(name, email, password, user_type, acct_created, last_login) VALUES(?, ?, ?, ?, ?, ?)`
var queryDeleteUser = `DELETE FROM users WHERE id = ?`
var queryBanUser = `UPDATE users SET banned = ? WHERE id=?`

// Categories Queries
var queryCateGetAll = `SELECT id, name FROM category ORDER BY id DESC`
var querySubCateGetAll = `SELECT id, name, parent_category FROM sub_category ORDER BY id DESC`
var queryCateAdd = `INSERT INTO category(name) VALUE (?)`
var querySubCateAdd = `INSERT INTO sub_category(name, parent_category) VALUE (?, ?)`
var queryCateDel = `DELETE FROM category WHERE id = ?`

// Sub-Categories Queries
