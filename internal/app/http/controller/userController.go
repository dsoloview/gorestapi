package controller

import (
	"encoding/json"
	"errors"
	"github.com/dsoloview/gorestapi/internal/app/http/response"
	"github.com/dsoloview/gorestapi/internal/app/http/validator"
	"github.com/dsoloview/gorestapi/internal/app/model"
	"github.com/dsoloview/gorestapi/internal/app/service"
	"io"
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
		if err == io.EOF {
			sendErrorResponse(w, errors.New("body is empty").Error())
		} else {
			sendErrorResponse(w, err.Error())
		}
		return
	}

	validationError := validator.NewUserValidator(&user).Validate()
	if validationError != nil {
		sendValidationError(w, validationError)
		return
	}

	createdUser, err := h.userService.CreateUser(&user)
	if err != nil {
		sendErrorResponse(w, err.Error())
		return
	}

	sendSuccessResponse(w, response.NewCreateUserResponse(createdUser))
}
