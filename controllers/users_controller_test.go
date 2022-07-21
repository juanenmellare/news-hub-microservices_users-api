package controllers

import (
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_users-api/services"
	"testing"
)

func TestNewUsersController(t *testing.T) {
	var userService services.UsersService

	assert.Implements(t, (*UsersController)(nil), NewUsersController(userService))
}
