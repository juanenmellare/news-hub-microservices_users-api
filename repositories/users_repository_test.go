package repositories

import (
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_users-api/databases"
	"testing"
)

func TestNewUsersRepository(t *testing.T) {
	var relationalDatabase databases.RelationalDatabase

	assert.Implements(t, (*UsersRepository)(nil), NewUsersRepository(relationalDatabase))
}
