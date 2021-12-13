package config

import (
	"github.com/golang-jwt/jwt"
)

type MyClaim struct {
	jwt.StandardClaims
	Username string
}

var mySigningKey = []byte("MySecrets")

func CreateToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "project-2",
		},
		Username: username,
	})

	signedStr, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}

	return signedStr
}
