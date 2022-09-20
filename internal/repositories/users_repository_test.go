package repositories

import (
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_users-api/internal/databases"
	"testing"
)

func Test_NewUsersRepository(t *testing.T) {
	var relationalDatabase databases.RelationalDatabase

	assert.Implements(t, (*UsersRepository)(nil), NewUsersRepository(relationalDatabase))
}
