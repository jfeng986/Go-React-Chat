package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(storedPassword, inputPassword string) error {
	log.Println(inputPassword, storedPassword)
	return bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(inputPassword))
}
