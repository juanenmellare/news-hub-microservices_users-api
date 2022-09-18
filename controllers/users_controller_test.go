package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"news-hub-microservices_users-api/api"
	"news-hub-microservices_users-api/mocks/models"
	mocks "news-hub-microservices_users-api/mocks/services"
	"news-hub-microservices_users-api/services"
	"testing"
)

func Test_NewUsersController(t *testing.T) {
	var userService services.UsersService

	assert.Implements(t, (*UsersController)(nil), NewUsersController(userService, "", 1))
}

func Test_usersControllerImpl_Create(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()

	userMock := models.NewUserBuilder().Build()

	var usersServiceMock mocks.UsersService
	usersServiceMock.On("Create", userMock.FirstName, userMock.LastName, userMock.Email, userMock.Password).
		Return(userMock.ID)

	controller := NewUsersController(&usersServiceMock, "", 1)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)
	reqBodyBytes := new(bytes.Buffer)
	request := &api.CreateUserRequest{
		FirstName:      &userMock.FirstName,
		LastName:       &userMock.LastName,
		Email:          &userMock.Email,
		Password:       &userMock.Password,
		PasswordRepeat: &userMock.Password,
	}

	_ = json.NewEncoder(reqBodyBytes).Encode(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/", reqBodyBytes)

	controller.Create(context)

	assert.Equal(t, fmt.Sprintf("{\"userId\":\"%s\"}", userMock.ID), writer.Body.String())
}

func Test_usersControllerImpl_Authenticate(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()

	userMock := models.NewUserBuilder().Build()

	var usersServiceMock mocks.UsersService
	usersServiceMock.On("Authenticate", userMock.Email, userMock.Password).Return(&userMock)

	controller := NewUsersController(&usersServiceMock, "", 1)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)
	reqBodyBytes := new(bytes.Buffer)
	request := &api.AuthenticateRequest{
		Email:    &userMock.Email,
		Password: &userMock.Password,
	}

	_ = json.NewEncoder(reqBodyBytes).Encode(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/login", reqBodyBytes)

	controller.Authenticate(context)

	assert.Contains(t, writer.Body.String(), "{\"token\":\"Bearer ")
}
