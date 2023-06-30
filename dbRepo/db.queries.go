package dbRepo

// Post Queries
var queryInsertPost = `INSERT INTO posts(title, content, user_id) VALUES(?,?,?)`
var queryGet3BlogPosts = `SELECT id, user_id, title, content FROM posts ORDER BY id DESC LIMIT ?`

// User Queries
var queryLoginUser = `SELECT id, password FROM users WHERE email=?`
var queryGetAllUsers = `SELECT id, name, email, user_type, banned FROM users ORDER BY id DESC`
var queryGetUserById = `SELECT name, email, password, acct_created, last_login, user_type, id FROM users WHERE id = ?`
var queryUpdateUser = `UPDATE users SET name=?, password=?, user_type=? WHERE id=?`
var queryFindByEmail = `SELECT email FROM users WHERE email = ?`
var queryAddUser = `INSERT INTO users(name, email, password, user_type, acct_created, last_login) VALUES(?, ?, ?, ?, ?, ?)`
var queryDeleteUser = `DELETE FROM users WHERE id = ?`
var queryBanUser = `UPDATE users SET banned = ? WHERE id=?`

// Categories Queries
var queryCateGetAll = `SELECT id, name FROM category ORDER BY id DESC`
var queryCateAdd = `INSERT INTO category(name) VALUE (?)`
var queryCategoryDelete = `DELETE FROM category WHERE id = ?`

// Sub-Categories Queries
var querySubCateGetAll = `SELECT id, name, parent_category FROM sub_category ORDER BY id DESC`
var queryGetSubCateById = `SELECT id, name, parent_category FROM sub_category WHERE parent_category = ?`
var querySubCateAdd = `INSERT INTO sub_category(name, parent_category) VALUE (?, ?)`
var queryDeleteSubCategory = `DELETE FROM sub_category WHERE id = ?`
var queryDeleteSubByParent = `DELETE FROM sub_category WHERE parent_category = ?`

// Menu Queries
var queryGetAllMenus = `SELECT id, name, is_navbar FROM menu ORDER BY id`
var queryCreateMenu = `INSERT INTO menu(name, is_navbar) VALUE (?,?)`
var queryEditIsNav = `UPDATE menu SET is_navbar=? WHERE id=?`
var queryFindIsNav = `SELECT id, name FROM menu where is_navbar = 1`
var queryDeleteMenu = `DELETE FROM menu WHERE id=?`
var queryCreateMenuLink = `INSERT INTO menu_item(name, target, parent_menu_id, position) VALUE(?,?,?,?)`
var queryDeleteMenuLink = `DELETE FROM menu_item WHERE id=?`
