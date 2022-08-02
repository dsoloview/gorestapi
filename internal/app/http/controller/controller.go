package controller

import (
	"encoding/json"
	"github.com/dsoloview/gorestapi/internal/app/http/response"
	"github.com/dsoloview/gorestapi/internal/app/service"
	"net/http"
	"net/url"
)

type Controller struct {
	UserController *UserController
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		UserController: NewUserController(service.User),
	}
}

func sendSuccessResponse(w http.ResponseWriter, response interface{}) {

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		sendErrorResponse(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func sendErrorResponse(w http.ResponseWriter, error string) {

	jsonResponse, err := json.Marshal(response.NewErrorResponse(error))

	if err != nil {
		sendErrorResponse(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write(jsonResponse)
}

func sendValidationError(w http.ResponseWriter, error url.Values) {

	result := map[string]interface{}{"validationError": error}
	jsonResponse, err := json.Marshal(result)
	if err != nil {
		sendErrorResponse(w, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write(jsonResponse)
}
