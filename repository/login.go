package repository

import (
	"github.com/nicolascancino/web-service-go/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (models.Usuario, bool) {
	user, founded, _ := CheckIfExistUser(email)

	if !founded {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return user, false
	}
	return user, true
}
