package repository

import "github.com/cwilliamson29/GoLangBlog/models"

type DatabaseRepo interface {
	InsertPost(newPost models.Post) error
	AuthenticateUser(email string, password string) (int, string, error)
	UpdateUser(u models.User) error
	GetUserById(id int) (models.User, error)
	GetBlogPost() (int, int, string, string, error)
	Get3BlogPost() (models.ArticleList, error)
}
