package utils

import (
	"code.google.com/p/go.crypto/bcrypt"
	"fmt"
	"github.com/robfig/revel"
)

func HashPassword(plainText string) (string, error) {
	bcryptCost, found := revel.Config.Int("user.bcrypt.cost")
	if !found {
		return plainText, fmt.Errorf("user.bcrypt.cost Error")
	}

	bytesPassword := []byte(plainText)
	bytesPassword, err := bcrypt.GenerateFromPassword(bytesPassword, bcryptCost)
	return string(bytesPassword), err
}
