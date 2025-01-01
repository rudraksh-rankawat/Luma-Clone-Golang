package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const hmacSampleSecret = "12qwaszx12"

func GenerateJWT(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": "",
		"email": "",
		"expiry":time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(hmacSampleSecret))

	return tokenString, err
}