package controller

import (
	"encoding/json"
	"github.com/dsoloview/gorestapi/internal/app/service"
	"net/http"
)

type Controller struct {
	UserController *UserController
}

type SuccessResponse struct {
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		UserController: NewUserController(service.User),
	}
}

func SendSuccessResponse(w http.ResponseWriter, response map[string]interface{}) {

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		SendErrorResponse(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func SendErrorResponse(w http.ResponseWriter, error string) {

	jsonResponse, err := json.Marshal(map[string]interface{}{
		"error": error,
	})

	if err != nil {
		SendErrorResponse(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write(jsonResponse)
}
