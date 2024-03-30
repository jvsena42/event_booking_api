package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "cc103de831561fc733db427e0bd1d319"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(secretKey)
}
