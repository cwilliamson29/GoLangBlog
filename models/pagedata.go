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
	UserAdd         map[string]interface{}
	UserDel         map[string]any
	CateAdd         map[string]any
	CateDel         map[string]any
	SubCateAdd      map[string]any
	SubCateDel      map[string]any
	Form            *forms.Form
	Data            map[int]interface{}
	Active          string
	IsAuthenticated int
}
