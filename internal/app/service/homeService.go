package service

type HomeService struct {
}

func NewHomeService() *HomeService {
	return &HomeService{}
}

func (h *HomeService) Hello() string {
	return "Hello"
}
