package rest

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_users-api/test/mocks/models"

	"testing"
)

func TestNewAuthenticateResponse(t *testing.T) {
	userMock := models.NewUserBuilder().Build()
	token := NewUserToken(1, &userMock)
	secretKey := ""
	response := NewAuthenticateResponse(token.ToString(secretKey))

	assert.Equal(t, fmt.Sprintf("Bearer %s", token.ToString(secretKey)), response.Token)
}

func TestNewCreateUserResponse(t *testing.T) {
	newUuid, _ := uuid.NewV4()
	response := NewCreateUserResponse(newUuid)

	assert.Equal(t, newUuid.String(), response.UserId)
}

func TestNewGetResponse(t *testing.T) {
	userMock := models.NewUserBuilder().Build()

	response := NewGetResponse(&userMock)

	assert.Equal(t, userMock.ID.String(), response.Id)
	assert.Equal(t, userMock.FirstName, response.FirstName)
	assert.Equal(t, userMock.LastName, response.LastName)
	assert.Equal(t, userMock.Email, response.Email)
}
