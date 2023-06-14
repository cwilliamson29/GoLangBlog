package dbrepo

import (
	"context"
	"errors"
	"github.com/cwilliamson29/GoLangBlog/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// Functions for accessing database

// InsertPost - Creating new a blog post
func (m *postgresDBRepo) InsertPost(newPost models.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `INSERT INTO posts(title, content, user_id) VALUES($1, $2, $3)`

	_, err := m.DB.ExecContext(ctx, query, newPost.Title, newPost.Content, newPost.UserID)
	if err != nil {
		return err
	}
	return nil
}

// GetUserById - Get a user from the database
func (m *postgresDBRepo) GetUserById(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//query := `SELECT name, email, password, user_type, id FROM users WHERE id = ?`
	query := `SELECT name, email, password, acct_created, last_login, user_type, id FROM users WHERE id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)

	var u models.User
	err := row.Scan(
		&u.Name,
		&u.Email,
		&u.Password,
		&u.AcctCreated,
		&u.LastLogin,
		&u.UserType,
		&u.ID,
	)
	if err != nil {
		log.Println(err)
		return u, err
	}
	return u, err
}

// AddUser - Addes a user to the database
func (m *postgresDBRepo) AddUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	findQuery := `SELECT email FROM users WHERE email = ?`
	row := m.DB.QueryRowContext(ctx, findQuery, u.Email)
	var emCheck interface{}
	err1 := row.Scan(&emCheck)

	query := `INSERT INTO users(name, email, password, user_type, acct_created, last_login) VALUES(?, ?, ?, ?, ?, ?)`
	hashedPassword, err2 := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err2 != nil {
		log.Println(err2)
	}
	log.Println("row: ", row)
	log.Println(row == nil)

	if err1 != nil {
		_, err := m.DB.ExecContext(ctx, query, u.Name, u.Email, hashedPassword, u.UserType, time.Now(), time.Now())
		if err != nil {
			return err
		}
		return nil
	} else {

		return errors.New("user already exists")
	}
	//log.Println("value of row: ", row)
	//
	//return nil
}

// UpdateUser - Updates a user in the database
func (m *postgresDBRepo) UpdateUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE users SET name=$1, email=$2, last_login=#3, user_type=$4`

	_, err := m.DB.ExecContext(ctx, query,
		u.Name,
		u.Email,
		time.Now(),
		u.UserType)

	if err != nil {
		return err
	}
	return nil
}

// AuthenticateUser - Checks database for user and logs in
func (m *postgresDBRepo) AuthenticateUser(email string, password string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id int
	var hashedPW string

	query := `SELECT id, password FROM users WHERE email=?`

	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(&id, &hashedPW)
	if err != nil {
		log.Println(err)
		return id, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPW), []byte(password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("password is incorrect")
	} else if err != nil {
		return 0, "", err
	}
	return id, hashedPW, nil
}

func (m *postgresDBRepo) GetBlogPost() (int, int, string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var id, uID int
	var aTitle, aContent string

	query := `SELECT id, user_id, title, content FROM posts LIMIT 1`

	row := m.DB.QueryRowContext(ctx, query)

	err := row.Scan(&id, &uID, &aTitle, &aContent)

	if err != nil {
		return id, uID, "", "", err
	}
	return id, uID, aTitle, aContent, nil
}

func (m *postgresDBRepo) Get3BlogPost() (map[int]interface{}, error) {
	var artList models.ArticleList
	artCollection := make(map[int]interface{})

	rows, err := m.DB.Query("SELECT id, user_id, title, content FROM posts ORDER BY id DESC LIMIT 3")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, uID int
		var title, content string
		err = rows.Scan(&id, &uID, &title, &content)
		if err != nil {
			panic(err)
		}

		artList.ID = id
		artList.UserID = uID
		artList.Title = title
		artList.Content = content
		artCollection[id] = artList
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return artCollection, nil
}

func (m *postgresDBRepo) GetAllUsers() (map[int]interface{}, error) {
	var user models.User
	userCollection := make(map[int]interface{})
	rows, err := m.DB.Query("SELECT name, email, user_type, id FROM users ORDER BY id DESC")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, uT int
		var name, email string
		err = rows.Scan(&name, &email, &uT, &id)
		if err != nil {
			panic(err)
		}

		user.ID = id
		user.Name = name
		user.Email = email
		user.UserType = uT
		userCollection[id] = user

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return userCollection, nil
}
