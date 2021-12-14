package config

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type MyClaim struct {
	jwt.StandardClaims
	Username string
}

var mySigningKey = []byte("MySecrets")

func CreateToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaim{
		StandardClaims: jwt.StandardClaims{},
		Username:       username,
	})

	signedStr, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}

	return signedStr
}

func VerifyToken(token string) bool {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return mySigningKey, nil
	})
	if err != nil {
		panic(err)
	}

	if _, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return true
	}

	return false
}
