package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken -
func CreateToken(id string, validated bool) (tokStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":       id,
		"exp":       time.Now().Add(time.Hour * 3).Unix(),
		"validated": validated,
	})
	return token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_KEY")))
}

// TokenValid -
func TokenValid(tokStr string) (*jwt.Token, error) {
	return jwt.Parse(tokStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_TOKEN_KEY")), nil
	})
}
