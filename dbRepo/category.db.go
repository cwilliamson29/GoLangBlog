package dbRepo

import (
	"errors"
	"github.com/cwilliamson29/GoLangBlog/models"
	"strconv"
)

// CateAdd - Creates category title
func (m *MySqlDB) CateAdd(n string) error {
	var success bool
	ct := m.Connect()
	if ct {
		success, _ = m.Insert(queryCateAdd, n)
	}
	// check if return true for success
	if success {
		return nil
	} else {
		return errors.New("Category not added")
	}
}

// SubCateAdd - Creates sub category title
func (m *MySqlDB) SubCateAdd(n string, id int) error {
	var success bool
	ct := m.Connect()
	if ct {
		success, _ = m.Insert(querySubCateAdd, n, id)
	}
	// check if return true for success
	if success {
		return nil
	} else {
		return errors.New("Sub category not added")
	}
}

// GetAllCategories - Gets a list of all users
func (m *MySqlDB) GetAllCategories() (map[int]interface{}, error) {
	var results *DbRow
	ct := m.Connect()
	if ct {
		results = m.Get(queryCateGetAll)
	}

	var cat models.Category
	cCollection := make(map[int]interface{})

	count := len(results.Row)
	c := count - 2

	for i := 0; i <= c; {
		id, _ := strconv.Atoi(results.Row[i])

		cat.ID = id
		cat.Name = results.Row[i+1]
		cCollection[id] = cat

		i = i + 2
	}
	return cCollection, nil
}

// GetAllSubCategories - Gets a list of all users
func (m *MySqlDB) GetAllSubCategories() (map[int]interface{}, error) {
	var results *DbRow
	ct := m.Connect()
	if ct {
		results = m.Get(querySubCateGetAll)
	}

	var cat models.SubCategory
	scCollection := make(map[int]interface{})

	count := len(results.Row)
	var c int
	if count >= 6 {
		c = count - 3
	} else {
		c = count
	}

	for i := 0; i <= c; {
		id, _ := strconv.Atoi(results.Row[i])
		pId, _ := strconv.Atoi(results.Row[i+2])

		cat.ID = id
		cat.Name = results.Row[i+1]
		cat.ParentCat = pId
		scCollection[id] = cat

		i = i + 3
	}
	return scCollection, nil
}
