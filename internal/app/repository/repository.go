package repository

import (
	"github.com/dsoloview/gorestapi/internal/app/model"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	User
}

type User interface {
	CreateUser(user *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{User: NewUserRepository(db)}
}
