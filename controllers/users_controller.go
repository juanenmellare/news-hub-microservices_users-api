package controllers

import "news-hub-microservices_users-api/services"

type UsersController interface {
}

type usersControllerImpl struct {
	userService services.UsersService
}

func NewUsersController(userService services.UsersService) UsersController {
	return &usersControllerImpl{
		userService,
	}
}
