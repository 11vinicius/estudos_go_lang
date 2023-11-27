package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password, passwordHashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
}
