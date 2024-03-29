package services

import (
	"fmt"
	"github.com/gofrs/uuid"
	"news-hub-microservices_users-api/internal/errors"
	"news-hub-microservices_users-api/internal/models"
	"news-hub-microservices_users-api/internal/repositories"
	"news-hub-microservices_users-api/internal/utils"
)

type UsersService interface {
	Create(firstName, lastName, email, password string) uuid.UUID
	Authenticate(email, password string) *models.User
	GetById(id string) *models.User
}

type usersService struct {
	usersRepository repositories.UsersRepository
	bCryptCost      int
}

func (u usersService) Create(firstName, lastName, email, password string) uuid.UUID {
	userFounded := u.usersRepository.FindByEmail(email)
	if userFounded != nil {
		panic(errors.NewAlreadyExistModelError(fmt.Sprintf("user with '%s' email", email)))
	}

	user := models.NewUser(firstName, lastName, email, password, u.bCryptCost)
	u.usersRepository.Create(user)

	return user.ID
}

func (u usersService) Authenticate(email, password string) *models.User {
	userFounded := u.usersRepository.FindByEmail(email)
	if userFounded == nil {
		panic(errors.NewInvalidEmailOrPasswordError())
	}

	doesPasswordMatch := utils.VerifyPassword(password, userFounded.Password)
	if !doesPasswordMatch {
		panic(errors.NewInvalidEmailOrPasswordError())
	}

	return userFounded
}

func (u usersService) GetById(id string) *models.User {
	return u.usersRepository.FindById(id)
}

func NewUsersService(usersRepository repositories.UsersRepository, bCryptCost int) UsersService {
	return &usersService{
		usersRepository,
		bCryptCost,
	}
}
