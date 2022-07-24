package service

import (
	"github.com/dsoloview/gorestapi/internal/app/helper"
	"github.com/dsoloview/gorestapi/internal/app/model"
	"github.com/dsoloview/gorestapi/internal/app/repository"
)

type UserService struct {
	repository repository.User
}

func NewUserService(repository repository.User) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
	encrypted, err := helper.EncryptPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.EncryptedPassword = encrypted
	createUser, err := s.repository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createUser, nil
}
