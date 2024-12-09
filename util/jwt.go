package util

import (
	"errors"
	"fmt"
	"os"

	"github.com/Jehanv60/helper"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey string

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	helper.GoDoEnv()
	SecretKey = os.Getenv("SecretKey")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return webtoken, nil
}

func VerifyToken(vertoken string) (*jwt.Token, error) {
	helper.GoDoEnv()
	SecretKey = os.Getenv("SecretKey")
	token, err := jwt.Parse(vertoken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("unexpected signing: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, errors.New(err.Error())
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
	return token.Header, nil
}
