package dbRepo

import "github.com/cwilliamson29/GoLangBlog/models"

type DatabaseRepo interface {
	InsertPost(newPost models.Post) error
	AuthenticateUser(email string, password string) (int, string, error)
	AddUser(u models.User) error
	UpdateUser(u models.User) error
	GetUserById(id int) (*models.User, error)
	GetBlogPost() (int, int, string, string, error)
	Get3BlogPost() (map[int]interface{}, error)
	GetAllUsers() (map[int]interface{}, error)
	DeleteUser(id int) error
	BanUser(id int, t int) error
	CateAdd(n string) error
	GetAllCategories() (map[int]interface{}, error)
	GetAllSubCategories() (map[int]interface{}, error)
	SubCateAdd(n string, id int) error
	CountSubCategoriesById(id int) (int, error)
	DeleteCategoryById(id int) error
	DeleteSubCategoryById(id int) error
	DeleteSubByParent(id int) error
}
