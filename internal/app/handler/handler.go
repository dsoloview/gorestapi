package handler

import (
	"github.com/dsoloview/gorestapi/internal/app/controller"
	"github.com/gorilla/mux"
)

type Handler struct {
	controllers *controller.Controller
	router      *mux.Router
}

func NewHandler(controllers *controller.Controller) *Handler {
	return &Handler{
		controllers: controllers,
		router:      mux.NewRouter(),
	}
}

func (h *Handler) Handle() *mux.Router {
	h.router.HandleFunc("/register", h.controllers.UserController.CreateUser).Methods("POST")

	return h.router
}
