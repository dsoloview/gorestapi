package service

import "github.com/dsoloview/gorestapi/internal/app/repository"

type Service struct {
	Home
}

type Home interface {
	Hello() string
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{Home: NewHomeService()}
}
