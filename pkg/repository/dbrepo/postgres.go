package dbrepo

import (
	"context"
	"errors"
	"github.com/cwilliamson29/GoLangBlog/models"
	"golang.org/x/crypto/bcrypt"
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

	query := `SELECT name, email, password, acct_created, last_login, user_type, id FROM users WHERE id = $1`

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
		return u, err
	}
	return u, err
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

	query := `SELECT id, password FROM users WHERE email=$1`

	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(&id, &hashedPW)

	if err != nil {
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
