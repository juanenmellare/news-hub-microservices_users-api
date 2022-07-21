package services

import "news-hub-microservices_users-api/repositories"

type UsersService interface {
}

type usersServiceImpl struct {
	usersRepository repositories.UsersRepository
}

func NewUsersService(usersRepository repositories.UsersRepository) UsersService {
	return &usersServiceImpl{
		usersRepository,
	}
}
