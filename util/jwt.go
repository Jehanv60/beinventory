package util

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var Secretkey string = "your-256-bit-secret"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(Secretkey))
	if err != nil {
		return "", err
	}
	return webtoken, nil
}

func VerifyToken(vertoken string) (*jwt.Token, error) {
	token, err := jwt.Parse(vertoken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("unexpected signing: %v", token.Header["alg"])
		}
		return []byte(Secretkey), nil
	})
	if err != nil {
		return nil, errors.New("error cookie")
	}
	return token, nil
}

func Decodetoken(vertoken string) (jwt.MapClaims, error) {
	token, err := VerifyToken(vertoken)
	if err != nil {
		return nil, err
	}
	claims, isok := token.Claims.(jwt.MapClaims)
	if isok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
