package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, cost int) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func VerifyPassword(candidatePassword, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword)); err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
