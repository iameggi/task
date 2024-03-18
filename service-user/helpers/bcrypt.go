package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash of the given password.
// It returns the hashed password as a byte slice.
func HashPassword(password []byte) string {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}

// ComparePassword compares the given hashed password with the provided password.
// It returns true if the passwords match, false otherwise.
func ComparePassword(hashedPsw, psw []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPsw, psw)
	return err == nil
}
