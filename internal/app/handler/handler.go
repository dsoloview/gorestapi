package handler

import (
	"github.com/dsoloview/gorestapi/internal/app/controller"
	"github.com/gorilla/mux"
)

type Handler struct {
	controllers *controller.Controller
}

func NewHandler(controllers *controller.Controller) *Handler {
	return &Handler{controllers: controllers}
}

func (h *Handler) Handle() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", h.controllers.Home.Index)

	return router
}
