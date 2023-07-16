package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(plainText string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(plainText), 10)
}

func CompareHashAndPlainText(hashed string, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainText))
	return err == nil
}
