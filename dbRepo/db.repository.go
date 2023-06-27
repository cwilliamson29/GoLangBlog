package dbRepo

import "github.com/cwilliamson29/GoLangBlog/models"

type DatabaseRepo interface {
	// User functions
	AuthenticateUser(email string, password string) (int, string, error)
	AddUser(u models.User) error
	UpdateUser(u models.User) bool
	GetUserById(id int) (*models.User, error)
	GetAllUsers() (map[int]interface{}, error)
	DeleteUser(id int) error
	BanUser(id int, t int) error
	// Blog post functions
	InsertPost(newPost models.Post) error
	GetBlogPost() (int, int, string, string, error)
	Get3BlogPost() (map[int]interface{}, error)
	// Category functions
	CateAdd(n string) error
	GetAllCategories() (map[int]interface{}, error)
	GetAllSubCategories() (map[int]interface{}, error)
	SubCateAdd(n string, id int) error
	CountSubCategoriesById(id int) (int, error)
	DeleteCategoryById(id int) error
	DeleteSubCategoryById(id int) error
	DeleteSubByParent(id int) error
	// Menu functions
	GetAllMenus() (map[int]interface{}, error)
	MenuCreate(n string, nav int) error
}
