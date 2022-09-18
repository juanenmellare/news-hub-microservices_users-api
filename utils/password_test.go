package utils

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword_Ok(t *testing.T) {
	password := "password"
	passwordHashed := HashPassword(password, bcrypt.MinCost)

	assert.NotNil(t, passwordHashed)
}

func TestHashPassword_Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("did not panic")
		} else {
			assert.Equal(t, bcrypt.InvalidCostError(32), r)
		}
	}()

	HashPassword("password", bcrypt.MaxCost+1)
}

func TestVerifyPassword_True(t *testing.T) {
	password := "password"
	passwordHashed := HashPassword(password, bcrypt.MinCost)

	assert.True(t, VerifyPassword(password, passwordHashed))
}

func TestVerifyPassword_False(t *testing.T) {
	password := "password"
	passwordHashed := HashPassword(password, bcrypt.MinCost)

	assert.False(t, VerifyPassword(password+"d", passwordHashed))
}
