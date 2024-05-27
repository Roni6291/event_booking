package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(pwd), err
}
