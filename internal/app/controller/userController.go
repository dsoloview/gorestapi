package controller

import (
	"fmt"
	"github.com/dsoloview/gorestapi/internal/app/service"
	"net/http"
)

type UserController struct {
	userService service.User
}

func NewHomeController(userService service.User) *UserController {
	return &UserController{userService: userService}
}

func (h *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
