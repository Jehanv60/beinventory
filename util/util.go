package util

import (
	"log"

	"github.com/Jehanv60/model/domain"
	"golang.org/x/crypto/bcrypt"
)

func Hashpassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func Unhashpassword(password, hashedpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	return err == nil
}

func ChangeMonth(countId []domain.Transaction) int {
	var sum int
	if countId != nil {
		for i := range countId {
			sum = 2 + i
			log.Println(sum, i)
		}
		return sum
	}
	return 1
}
