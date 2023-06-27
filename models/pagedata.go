package models

import (
	"github.com/cwilliamson29/GoLangBlog/forms"
)

type PageData struct {
	UserInfo        map[string]string
	StrMap          map[string]string
	IntMap          map[string]int
	FltMap          map[string]float32
	DataMap         map[string]interface{}
	Warning         string
	Error           error
	Success         string
	Status          map[string]any // For updating form submission status
	Form            *forms.Form
	Data            map[int]interface{}
	Data2           map[int]interface{}
	Active          string
	UA              string
	CA              string
	MA              string
	IsAuthenticated int
}
