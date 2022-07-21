package models

import (
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, emptyString, user.Salt)
}
