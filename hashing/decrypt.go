package hashing

import "golang.org/x/crypto/bcrypt"

func CheckPassword(hashedPassword []byte, enteredPassword string) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(enteredPassword))
	return err == nil
}
