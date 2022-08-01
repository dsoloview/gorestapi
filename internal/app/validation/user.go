package validation

import (
	"github.com/dsoloview/gorestapi/internal/app/model"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

type UserValidation struct {
	user model.User
}

func NewUserValidation(user model.User) *UserValidation {
	return &UserValidation{user: user}
}

func (v *UserValidation) rules() govalidator.MapData {
	return govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required", "min:6"},
	}
}

func (v *UserValidation) messages() govalidator.MapData {
	return govalidator.MapData{
		"email":    []string{"required:Email is required", "email:Incorrect email"},
		"password": []string{"required:Password is required", "min:Password must be longer than 6 characters"},
	}
}

func (v *UserValidation) Validate() url.Values {

	opts := govalidator.Options{
		Data:     &v.user,
		Rules:    v.rules(),
		Messages: v.messages(),
	}

	validator := govalidator.New(opts)
	e := validator.ValidateStruct()

	if len(e) > 0 {
		return e
	}
	return nil

}
