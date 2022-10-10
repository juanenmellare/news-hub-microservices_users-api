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

	fmt.Println("1")

	user := models.NewUser(firstName, lastName, email, password, u.bCryptCost)

	fmt.Println("2")

	u.usersRepository.Create(user)

	fmt.Println("3")

	return user.ID
}

func (u usersService) Authenticate(email, password string) *models.User {
	fmt.Println("1")
	userFounded := u.usersRepository.FindByEmail(email)
	if userFounded == nil {
		panic(errors.NewInvalidEmailOrPasswordError())
	}
	fmt.Println("2")

	doesPasswordMatch := utils.VerifyPassword(password, userFounded.Password)
	fmt.Println("3")
	if !doesPasswordMatch {
		panic(errors.NewInvalidEmailOrPasswordError())
	}
	fmt.Println("4")

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
