package service

import (
	"github.com/dsoloview/gorestapi/internal/app/model"
	"github.com/dsoloview/gorestapi/internal/app/repository"
)

type Service struct {
	User
}

type User interface {
	CreateUser(user *model.User) (*model.User, error)
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{User: NewUserService(repositories.User)}
}
