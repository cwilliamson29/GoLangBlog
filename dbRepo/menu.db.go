package dbRepo

import (
	"errors"
	"github.com/cwilliamson29/GoLangBlog/models"
	"log"
	"strconv"
)

// GetAllMenus - Gets a list of all users
func (m *MySqlDB) GetAllMenus() (map[int]interface{}, error) {
	var results *DbRow
	ct := m.Connect()
	if ct {
		results = m.Get(queryGetAllMenus)
	}

	var menu models.MainMenu
	Collection := make(map[int]interface{})

	count := len(results.Row)
	c := count - 3

	for i := 0; i <= c; {
		id, _ := strconv.Atoi(results.Row[i])
		isn, _ := strconv.Atoi(results.Row[i+2])

		menu.ID = id
		menu.Name = results.Row[i+1]
		menu.IsNavbar = isn
		Collection[id] = menu

		i = i + 3
	}
	return Collection, nil
}

// MenuCreate - Creates menu
func (m *MySqlDB) MenuCreate(n string, nav int) error {
	var success bool
	ct := m.Connect()
	if ct {
		results, err := m.IsNavFind()
		if err != nil {
			log.Println(err)
		}
		c := len(results)
		if nav == 1 && c > 0 {
			return errors.New("Main navbar already exists")
		} else {
			success, _ = m.Insert(queryCreateMenu, n, nav)
		}
	}
	// check if return true for success
	if success {
		return nil
	} else {
		return errors.New("Sub category not added")
	}
}

// IsNavFind - finds entry if 'is_navbar' = true
func (m *MySqlDB) IsNavFind() ([]string, error) {
	var results *DbRow
	ct := m.Connect()
	if ct {
		results = m.Get(queryFindIsNav)
	}
	rtn := results.Row
	// check if return true for success
	return rtn, nil
}

// UpdateUser - Updates a user in the database
func (m *MySqlDB) UpdateIsNav(n int, id int) bool {
	var suc bool
	ct := m.Connect()
	if ct {
		suc = m.Update(queryEditIsNav, n, id)
	}
	return suc
}

// DeleteCategoryById
func (m *MySqlDB) DeleteMenuById(id int) error {
	var success bool
	ct := m.Connect()
	if ct {
		success = m.Delete(queryDeleteMenu, id)
	}
	// check if return true for success
	if success {
		return nil
	} else {
		return errors.New("Menu has references to menu items")
	}
}
