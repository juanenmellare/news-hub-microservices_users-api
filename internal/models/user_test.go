package models

import (
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_users-api/internal/utils"
	"testing"
)

func TestUser(t *testing.T) {
	user := &User{}

	var emptyString string
	assert.Equal(t, uuid.UUID{}, user.ID)
	assert.Equal(t, emptyString, user.FirstName)
	assert.Equal(t, emptyString, user.LastName)
	assert.Equal(t, emptyString, user.Email)
	assert.Equal(t, emptyString, user.Password)
}

func TestNewUser(t *testing.T) {
	firstName := "firstName"
	lastName := "lastName"
	email := "email"
	password := "password"
	cost := 10

	user := NewUser(firstName, lastName, email, password, cost)

	assert.Equal(t, firstName, user.FirstName)
	assert.Equal(t, lastName, user.LastName)
	assert.Equal(t, email, user.Email)
	assert.True(t, utils.VerifyPassword(password, user.Password))
}
