package config

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/nicolascancino/web-service-go/models"
)

func GenerateJWT(user *models.Usuario) (string, error) {

	myKey := []byte("CLAVE_JWT")
	payload := jwt.MapClaims{
		"email":     user.Email,
		"nombre":    user.Nombre,
		"apellidos": user.Apellidos,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
