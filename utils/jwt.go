package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var privateKey *ecdsa.PrivateKey

func init() {

	keyFile := "keyfile.pem"

	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			panic(err)
		}

		keyBytes, err := x509.MarshalECPrivateKey(privateKey)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(keyFile, keyBytes, 0600)
		if err != nil {
			panic(err)
		}
	} else {
		keyBytes, err := os.ReadFile(keyFile)
		if err != nil {
			panic(err)
		}

		privateKey, err = x509.ParseECPrivateKey(keyBytes)
		if err != nil {
			panic(err)
		}
	}
}

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(privateKey)
}

func VerifyToken(token string) (int64, error) {

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodECDSA)

		if !ok {
			log.Println("ERROR VerifyToken: Unexpected signing method")
			return nil, errors.New("Unexpected signing method")
		}

		return &privateKey.PublicKey, nil
	})

	if err != nil {
		log.Println("ERROR VerifyToken: ", err)
		return 0, errors.New("Could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		log.Println("ERROR VerifyToken: Invalid Token")
		return 0, errors.New("Invalid Token!")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		log.Println("ERROR VerifyToken: Invalid Token claims")
		return 0, errors.New("Invalid token claims!")
	}

	//email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
