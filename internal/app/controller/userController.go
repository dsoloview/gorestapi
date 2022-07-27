package controller

import (
	"encoding/json"
	"github.com/dsoloview/gorestapi/internal/app/model"
	"github.com/dsoloview/gorestapi/internal/app/response"
	"github.com/dsoloview/gorestapi/internal/app/service"
	"net/http"
)

type UserController struct {
	userService service.User
}

func NewUserController(userService service.User) *UserController {
	return &UserController{userService: userService}
}

func (h *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendErrorResponse(w, err.Error())
		return
	}
	createdUser, err := h.userService.CreateUser(&user)
	if err != nil {
		sendErrorResponse(w, err.Error())
		return
	}

	sendSuccessResponse(w, response.NewCreateUserResponse(createdUser))
}
