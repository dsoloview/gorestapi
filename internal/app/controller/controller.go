package controller

import (
	"github.com/dsoloview/gorestapi/internal/app/service"
	"net/http"
)

type Controller struct {
	service *service.Service
	Home
}

type Home interface {
	Index(w http.ResponseWriter, r *http.Request)
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		service: service,
		Home:    NewHomeController(service.Home),
	}
}
