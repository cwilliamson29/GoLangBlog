package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) HasRequired(tagIDs ...string) {
	for _, tagID := range tagIDs {
		value := f.Get(tagID)
		if strings.TrimSpace(value) == "" {
			f.Errors.AddError(tagID, "This field can't be blank")
		}
	}
}

func (f *Form) HasValue(tagID string, r *http.Request) bool {
	x := r.Form.Get(tagID)
	//if x == "" {
	//	f.Errors.AddError(tagID, "Field Empty")
	//	return false
	//} else {
	//	return true
	//}
	return x != ""
}

func (f *Form) MinLength(tagID string, length int, r *http.Request) bool {
	x := r.Form.Get(tagID)
	if len(x) < length {
		f.Errors.AddError(tagID, fmt.Sprintf("This field mud be %d characters long or more", length))
		return false
	}
	return true
}

func (f *Form) IsEmail(tagID string) {
	if !govalidator.IsEmail(f.Get(tagID)) {
		f.Errors.AddError(tagID, "Invalid Email")
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
