package handler

import (
	"fmt"
	"github.com/dsoloview/gorestapi/internal/app/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Handle() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", h.HomeHandler)

	return router
}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Category: %v\n", 2)
	if err != nil {
		logrus.Debug(err)
	}
}
