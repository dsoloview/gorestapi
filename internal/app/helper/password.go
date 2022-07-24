package helper

import (
	"github.com/dsoloview/gorestapi/internal/app/model"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", nil
	}
	return string(encrypted), nil
}

func ComparePasswords(password string, user *model.User) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
