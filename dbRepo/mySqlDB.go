package dbRepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/cwilliamson29/GoLangBlog/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
)

// Functions for accessing database

// InsertPost - Creating new a blog post
func (m *MySqlDB) InsertPost(newPost models.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//query := `INSERT INTO posts(title, content, user_id) VALUES($1, $2, $3)`

	_, err := m.DB.ExecContext(ctx, queryInsertPost, newPost.Title, newPost.Content, newPost.UserID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetUserById - Get a user from the database
func (m *MySqlDB) GetUserById(id int) (*models.User, error) {
	var results *DbRow
	ct := m.Connect()
	if ct {
		results = m.Get(queryGetUserById, id)
	}

	var u models.User
	ac, _ := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", results.Row[3])
	ll, _ := time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", results.Row[4])
	ut, _ := strconv.Atoi(results.Row[5])
	uId, _ := strconv.Atoi(results.Row[6])

	u.Name = results.Row[0]
	u.Email = results.Row[1]
	u.Password = results.Row[2]
	u.AcctCreated = ac
	u.LastLogin = ll
	u.UserType = ut
	u.ID = uId

	return &u, nil
}

// AddUser - Addes a user to the database
func (m *MySqlDB) AddUser(u models.User) error {
	var suc bool
	var id int64

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	ct := m.Connect()
	if ct {
		suc, id = m.Insert(queryAddUser, u.Name, u.Email, hashedPassword, u.UserType, time.Now(), time.Now())
		if !suc {
			er := fmt.Sprintf("User not added: ", id)
			return errors.New(er)
		}
	}
	return nil
}

// UpdateUser - Updates a user in the database
func (m *MySqlDB) UpdateUser(u models.User) error {
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
func (m *MySqlDB) AuthenticateUser(email string, password string) (int, string, error) {
	var results *DbRow
	ct := m.Connect()
	if ct {
		results = m.Get(queryLoginUser, email)
	}

	id, _ := strconv.Atoi(results.Row[0])
	hashedPW := results.Row[1]

	err := bcrypt.CompareHashAndPassword([]byte(hashedPW), []byte(password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("password is incorrect")
	} else if err != nil {
		return 0, "", err
	}
	return id, hashedPW, nil
}

func (m *MySqlDB) GetBlogPost() (int, int, string, string, error) {
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

// Get3BlogPost - Gets first 3 blog posts out of DB
func (m *MySqlDB) Get3BlogPost() (map[int]interface{}, error) {
	var results *DbRow
	ct := m.Connect()
	if ct {
		results = m.Get(queryGet3BlogPosts, 3)
	}

	var artList models.ArticleList
	artCollection := make(map[int]interface{})

	for i := 0; i <= 8; {
		id, _ := strconv.Atoi(results.Row[i])
		uId, _ := strconv.Atoi(results.Row[i+1])

		artList.ID = id
		artList.UserID = uId
		artList.Title = results.Row[i+2]
		artList.Content = results.Row[i+3]
		artCollection[id] = artList

		i = i + 4
	}

	return artCollection, nil
}

// GetAllUsers - Gets a list of all users
func (m *MySqlDB) GetAllUsers() (map[int]interface{}, error) {
	var results *DbRow
	ct := m.Connect()
	if ct {
		results = m.Get(queryGetAllUsers)
	}

	var user models.User
	userCollection := make(map[int]interface{})

	count := len(results.Row)
	c := count - 4

	for i := 0; i <= c; {
		id, _ := strconv.Atoi(results.Row[i])
		uT, _ := strconv.Atoi(results.Row[i+3])

		user.ID = id
		user.Name = results.Row[i+1]
		user.Email = results.Row[i+2]
		user.UserType = uT
		userCollection[id] = user

		i = i + 4
	}
	return userCollection, nil
}

// DeleteUser - Deletes user from system
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

// BanUser - Bans user from further comments
func (m *MySqlDB) BanUser(id int) error {
	var success bool
	ct := m.Connect()
	if ct {
		success = m.Update(queryBanUser, id)
	}
	// check if return true for success
	if success {
		return nil
	} else {
		return errors.New("user not banned")
	}
}
