package api

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_users-api/mocks/models"
	"testing"
)

func TestNewAuthenticateResponse(t *testing.T) {
	userMock := models.NewUserBuilder().Build()
	token := NewUserToken(1, &userMock)
	secretKey := ""
	response := NewAuthenticateResponse(token, secretKey)

	assert.Equal(t, fmt.Sprintf("Bearer %s", token.ToString(secretKey)), response.Token)
}

func TestNewCreateUserResponse(t *testing.T) {
	newUuid, _ := uuid.NewV4()
	response := NewCreateUserResponse(newUuid)

	assert.Equal(t, newUuid.String(), response.UserId)
}
