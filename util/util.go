package util

import "golang.org/x/crypto/bcrypt"

func Hashpassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func Unhashpassword(password, hashedpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	return err == nil
}
