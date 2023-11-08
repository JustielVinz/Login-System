package hashing

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	cost := bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return hashedPassword, err
}
