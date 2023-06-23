package models

import "time"

type User struct {
	Name        string
	Email       string
	Password    string
	AcctCreated time.Time
	LastLogin   time.Time
	UserType    int
	Banned      int
	ID          int
}

type Post struct {
	Title   string
	Content string
	UserID  int
	ID      int
}

type Category struct {
	Name string
	ID   int
}
