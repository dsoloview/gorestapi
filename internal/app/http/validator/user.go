package validator

import (
	"github.com/dsoloview/gorestapi/internal/app/model"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

type UserValidator struct {
	user *model.User
}

func NewUserValidator(user *model.User) *UserValidator {
	return &UserValidator{user: user}
}

func (v *UserValidator) rules() govalidator.MapData {
	return govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required", "min:6"},
	}
}

func (v *UserValidator) messages() govalidator.MapData {
	return govalidator.MapData{
		"email":    []string{"required:Email is required", "email:Incorrect email"},
		"password": []string{"required:Password is required", "min:Password must be longer than 6 characters"},
	}
}

func (v *UserValidator) Validate() url.Values {

	opts := govalidator.Options{
		Data:     v.user,
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
