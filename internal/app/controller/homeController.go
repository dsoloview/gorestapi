package controller

import (
	"fmt"
	"github.com/dsoloview/gorestapi/internal/app/service"
	"github.com/sirupsen/logrus"
	"net/http"
)

type HomeController struct {
	homeService service.Home
}

func NewHomeController(homeService service.Home) *HomeController {
	return &HomeController{homeService: homeService}
}

func (h *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, h.homeService.Hello())
	if err != nil {
		logrus.Debug(err)
	}
}
