package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

	assert.Implements(t, (*UsersController)(nil), NewUsersController(userService))
}

func Test_usersControllerImpl_Create(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()

	var usersServices mocks.UsersService
	usersServices.On("Create", mock.AnythingOfType("*models.User"))

	service := NewUsersController(&usersServices)

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)
	reqBodyBytes := new(bytes.Buffer)

	userMock := models.NewUserBuilder().Build()

	createUserRequest := &api.CreateUserRequest{
		FirstName:      &userMock.FirstName,
		LastName:       &userMock.LastName,
		Email:          &userMock.Email,
		Password:       &userMock.Password,
		PasswordRepeat: &userMock.Password,
	}

	_ = json.NewEncoder(reqBodyBytes).Encode(createUserRequest)
	context.Request, _ = http.NewRequest(http.MethodPost, "/", reqBodyBytes)

	service.Create(context)

	assert.Equal(t, "{\"id\":\"00000000-0000-0000-0000-000000000000\",\"firstName\":\"foo-firstname\",\"lastName\":\"foo-lastname\",\"email\":\"foo-email@email.com\",\"password\":\"password\",\"salt\":\"10\"}", writer.Body.String())
}
