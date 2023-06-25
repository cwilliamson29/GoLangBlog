package dbRepo

import (
	"errors"
	"fmt"
	"github.com/cwilliamson29/GoLangBlog/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
)

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

// AddUser - Adds a user to the database
func (m *MySqlDB) AddUser(u models.User) error {
	var exists *DbRow
	var suc bool
	var id int64
	ct := m.Connect()
	if ct {
		exists = m.Get(queryFindByEmail, u.Email)
		if len(exists.Row) != 0 {
			er := fmt.Sprintf("User not added: %s, already exists", exists.Row[0])
			return errors.New(er)
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

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
func (m *MySqlDB) UpdateUser(u models.User) bool {
	var suc bool
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	ct := m.Connect()
	if ct {
		suc = m.Update(queryUpdateUser, u.Name, hashedPassword, u.UserType, u.ID)
	}
	return suc
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
	c := count - 5

	for i := 0; i <= c; {
		id, _ := strconv.Atoi(results.Row[i])
		uT, _ := strconv.Atoi(results.Row[i+3])
		b, _ := strconv.Atoi(results.Row[i+4])

		user.ID = id
		user.Name = results.Row[i+1]
		user.Email = results.Row[i+2]
		user.UserType = uT
		user.Banned = b
		userCollection[id] = user

		i = i + 5
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
func (m *MySqlDB) BanUser(id int, t int) error {
	var success bool
	ct := m.Connect()
	if ct {
		success = m.Update(queryBanUser, t, id)
	}
	// check if return true for success
	if success {
		return nil
	} else {
		return errors.New("user not banned")
	}
}
