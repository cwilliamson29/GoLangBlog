package dbRepo

import (
	"context"
	"github.com/cwilliamson29/GoLangBlog/models"
	"log"
	"strconv"
	"time"
)

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
