package validator

import (
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

type Validator interface {
	rules() govalidator.MapData
	messages() govalidator.MapData
	Validate() url.Values
}
