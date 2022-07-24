package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 1 {
		enc, err := encryptPassword(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

func encryptPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", nil
	}
	return string(result), nil
}

func comparePasswords(password string, user *User) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
