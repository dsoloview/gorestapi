package controller

import (
	"github.com/dsoloview/gorestapi/internal/app/service"
)

type Controller struct {
	service        *service.Service
	UserController *UserController
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		service:        service,
		UserController: NewHomeController(service.User),
	}
}
