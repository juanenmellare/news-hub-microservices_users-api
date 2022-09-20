package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"news-hub-microservices_users-api/internal/errors"
	"news-hub-microservices_users-api/internal/rest"
	"news-hub-microservices_users-api/internal/services"
	modelsMocks "news-hub-microservices_users-api/test/mocks/models"
	servicesMocks "news-hub-microservices_users-api/test/mocks/services"
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

	userMock := modelsMocks.NewUserBuilder().Build()

	var usersServiceMock servicesMocks.UsersService
	usersServiceMock.On("Create", userMock.FirstName, userMock.LastName, userMock.Email, userMock.Password).
		Return(userMock.ID)

	controller := NewUsersController(&usersServiceMock, "", 1)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)
	reqBodyBytes := new(bytes.Buffer)
	request := &rest.CreateUserRequest{
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

	userMock := modelsMocks.NewUserBuilder().Build()

	var usersServiceMock servicesMocks.UsersService
	usersServiceMock.On("Authenticate", userMock.Email, userMock.Password).Return(&userMock)

	controller := NewUsersController(&usersServiceMock, "", 1)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)
	reqBodyBytes := new(bytes.Buffer)
	request := &rest.AuthenticateRequest{
		Email:    &userMock.Email,
		Password: &userMock.Password,
	}

	_ = json.NewEncoder(reqBodyBytes).Encode(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/login", reqBodyBytes)

	controller.Authenticate(context)

	assert.Contains(t, writer.Body.String(), "{\"token\":\"Bearer ")
}

func Test_usersControllerImpl_Get(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()

	userMock := modelsMocks.NewUserBuilder().Build()

	var usersServiceMock servicesMocks.UsersService
	usersServiceMock.On("GetById", userMock.ID.String()).Return(&userMock)

	secret := "foo"
	controller := NewUsersController(&usersServiceMock, secret, 1)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	token := rest.NewUserToken(1, &userMock)
	tokenString := token.ToString(secret)

	context.Request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenString))

	controller.Get(context)

	assert.Equal(t, "{\"firstName\":\"foo-firstname\",\"lastName\":\"foo-lastname\",\"email\":\"foo-email@email.com\"}", writer.Body.String())
}

func Test_usersControllerImpl_Get_not_found_user_from_token(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, &errors.ApiError{Code: 404, Status: "Not Found", Message: "user from token not found"}, r)
		} else {
			t.Errorf("did not panic")
		}
	}()

	userMock := modelsMocks.NewUserBuilder().Build()

	var usersServiceMock servicesMocks.UsersService
	usersServiceMock.On("GetById", userMock.ID.String()).Return(nil)

	secret := "foo"
	controller := NewUsersController(&usersServiceMock, secret, 1)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
	token := rest.NewUserToken(1, &userMock)
	tokenString := token.ToString(secret)

	context.Request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenString))

	controller.Get(context)
}
