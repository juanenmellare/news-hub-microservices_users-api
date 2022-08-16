package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_users-api/errors"
	"news-hub-microservices_users-api/mocks/models"
	mocks "news-hub-microservices_users-api/mocks/repositories"
	"testing"
)

func Test_NewUsersService(t *testing.T) {
	var userRepository mocks.UsersRepository

	assert.Implements(t, (*UsersService)(nil), NewUsersService(&userRepository))
}

func Test_usersServiceImpl_Create(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()

	userMock := models.NewUserBuilder().Build()

	var userRepository mocks.UsersRepository
	userRepository.On("FindByEmail", userMock.Email).Return(nil)
	userRepository.On("Create", &userMock)

	service := NewUsersService(&userRepository)

	service.Create(&userMock)
}

func Test_usersServiceImpl_Create_NewAlreadyExistModelError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, &errors.AlreadyExistModelError{Message: "user with 'foo-email@email.com' email already exist"}, r)
		} else {
			panic("should panic!")
		}
	}()

	userMock := models.NewUserBuilder().Build()

	var userRepository mocks.UsersRepository
	userRepository.On("FindByEmail", userMock.Email).Return(&userMock)

	service := NewUsersService(&userRepository)

	service.Create(&userMock)
}
