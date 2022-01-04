package tool

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

type MyClaim struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func TokenCreate(id int) string {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaim{
		ID:             id,
		StandardClaims: jwt.StandardClaims{},
	})

	signedStr, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return signedStr
}

func TokenValidate(t string) (*jwt.Token, error) {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	token, _ := jwt.ParseWithClaims(t, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error while decoding the token")
		}

		return jwtKey, nil
	})

	return token, nil
}
