package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const PassWordCost = 12

func SetPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return "", err
	}
	passwordDigest := string(bytes)
	return passwordDigest, nil
}

func CheckPassword(password string, passwordDigest string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(passwordDigest))

	return err == nil
}
