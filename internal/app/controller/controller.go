package controller

import (
	"github.com/dsoloview/gorestapi/internal/app/service"
)

type Controller struct {
	UserController *UserController
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		UserController: NewUserController(service.User),
	}
}
