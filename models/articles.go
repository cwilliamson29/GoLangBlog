package models

type Article struct {
	BlogTitle   string
	BlogArticle string
	ID          []int
	UserID      []int
}

type ArticleList struct {
	ID      []int
	UserID  []int
	Title   []string
	Content []string
}
