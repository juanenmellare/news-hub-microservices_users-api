package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"news-hub-microservices_users-api/internal/errors"
	"news-hub-microservices_users-api/internal/models"
	mocksBuilders "news-hub-microservices_users-api/test/mocks/models"
	"news-hub-microservices_users-api/test/mocks/repositories"
	"testing"
)

var bCryptCost = bcrypt.DefaultCost

func Test_NewUsersService(t *testing.T) {
	var userRepository mocks.UsersRepository

	assert.Implements(t, (*UsersService)(nil), NewUsersService(&userRepository, bCryptCost))
}
func assertShouldNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()
}

func Test_usersServiceImpl_Create(t *testing.T) {
	assertShouldNotPanic(t)

	userMock := mocksBuilders.NewUserBuilder().Build()

	var userRepository mocks.UsersRepository
	userRepository.On("FindByEmail", userMock.Email).Return(nil)
	userRepository.On("Create", mock.AnythingOfType("*models.User")).Run(func(args mock.Arguments) {
		user := args.Get(0).(*models.User)
		user.ID = userMock.ID
	})

	service := NewUsersService(&userRepository, bCryptCost)

	userId := service.Create(userMock.FirstName, userMock.LastName, userMock.Email, userMock.Password)

	assert.Equal(t, userMock.ID, userId)
}

func Test_usersServiceImpl_Create_NewAlreadyExistModelError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, &errors.AlreadyExistModelError{Message: "user with 'foo-email@email.com' email already exist"}, r)
		} else {
			panic("did not panic")
		}
	}()

	userMock := mocksBuilders.NewUserBuilder().Build()

	var userRepository mocks.UsersRepository
	userRepository.On("FindByEmail", userMock.Email).Return(&userMock)

	service := NewUsersService(&userRepository, bCryptCost)

	service.Create(userMock.FirstName, userMock.LastName, userMock.Email, userMock.Password)
}

func Test_usersServiceImpl_Authenticate(t *testing.T) {
	assertShouldNotPanic(t)

	userMock := mocksBuilders.NewUserBuilder().Build()

	var userRepository mocks.UsersRepository
	userRepository.On("FindByEmail", userMock.Email).Return(&userMock)

	service := NewUsersService(&userRepository, bCryptCost)

	user := service.Authenticate(userMock.Email, "password")

	assert.Equal(t, userMock.ID, user.ID)
}

func Test_usersServiceImpl_Authenticate_NewInvalidEmailOrPasswordError_FindByEmail(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, &errors.InvalidEmailOrPasswordError{Message: "Invalid Email or Password"}, r)
		} else {
			panic("did not panic")
		}
	}()

	userMock := mocksBuilders.NewUserBuilder().Build()

	var userRepository mocks.UsersRepository
	userRepository.On("FindByEmail", userMock.Email).Return(nil)

	service := NewUsersService(&userRepository, bCryptCost)

	_ = service.Authenticate(userMock.Email, "")
}

func Test_usersServiceImpl_Authenticate_NewInvalidEmailOrPasswordError_VerifyPassword(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, &errors.InvalidEmailOrPasswordError{Message: "Invalid Email or Password"}, r)
		} else {
			panic("did not panic")
		}
	}()

	userMock := mocksBuilders.NewUserBuilder().Build()

	var userRepository mocks.UsersRepository
	userRepository.On("FindByEmail", userMock.Email).Return(&userMock)

	service := NewUsersService(&userRepository, bCryptCost)

	_ = service.Authenticate(userMock.Email, "")
}

func Test_usersServiceImpl_GetById(t *testing.T) {
	userMock := mocksBuilders.NewUserBuilder().Build()

	var userRepository mocks.UsersRepository
	userRepository.On("FindById", userMock.ID.String()).Return(&userMock)

	service := NewUsersService(&userRepository, bCryptCost)

	user := service.GetById(userMock.ID.String())

	assert.Equal(t, userMock.ID, user.ID)
}
