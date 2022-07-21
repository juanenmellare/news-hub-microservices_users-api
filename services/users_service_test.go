package services

import (
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_users-api/repositories"

	"testing"
)

func TestNewUsersService(t *testing.T) {
	var userRepository repositories.UsersRepository

	assert.Implements(t, (*UsersService)(nil), NewUsersService(userRepository))
}
