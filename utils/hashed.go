package utils

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func Hashed(password string) string {
	rounds, _ := strconv.Atoi(os.Getenv("JWT_SECRET"))
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), rounds)
	return string(hash)
}
