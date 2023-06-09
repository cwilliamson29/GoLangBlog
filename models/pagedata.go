package models

import (
	"github.com/cwilliamson29/GoLangBlog/pkg/forms"
)

type PageData struct {
	UserInfo        map[string]string
	StrMap          map[string]string
	IntMap          map[string]int
	FltMap          map[string]float32
	DataMap         map[string]interface{}
	CSRFToken       string
	Warning         string
	Error           string
	Form            *forms.Form
	Data            map[string]interface{}
	IsAuthenticated int
}
