package authentication

import (
	"golang.org/x/crypto/bcrypt"
)

func CreatePasswordHash(plainPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
}

func ComparePsswordHashToPlainPassword(passwordHash []byte, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword(passwordHash, []byte(plainPassword))

	return err == nil
}
