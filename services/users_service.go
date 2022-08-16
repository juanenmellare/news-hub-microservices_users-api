package services

import (
	"fmt"
	"news-hub-microservices_users-api/errors"
	"news-hub-microservices_users-api/models"
	"news-hub-microservices_users-api/repositories"
)

type UsersService interface {
	Create(user *models.User)
}

type usersServiceImpl struct {
	usersRepository repositories.UsersRepository
}

func (u usersServiceImpl) Create(user *models.User) {
	email := user.Email
	userFounded := u.usersRepository.FindByEmail(email)
	if userFounded != nil {
		panic(errors.NewAlreadyExistModelError(fmt.Sprintf("user with '%s' email", email)))
	}
	u.usersRepository.Create(user)
}

func NewUsersService(usersRepository repositories.UsersRepository) UsersService {
	return &usersServiceImpl{
		usersRepository,
	}
}
