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
	c := count - 2

	for i := 0; i <= c; {
		id, _ := strconv.Atoi(results.Row[i])

		menu.ID = id
		menu.Name = results.Row[i+1]
		Collection[id] = menu

		i = i + 2
	}
	return Collection, nil
}

// MenuCreate - Creates sub category title
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
