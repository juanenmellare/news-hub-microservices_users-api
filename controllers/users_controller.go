package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news-hub-microservices_users-api/api"
	"news-hub-microservices_users-api/services"
)

type UsersController interface {
	Create(context *gin.Context)
}

type usersControllerImpl struct {
	userService services.UsersService
}

func (u usersControllerImpl) Create(context *gin.Context) {
	var createUserRequest api.CreateUserRequest
	createUserRequest.MarshallAndValidate(context)

	context.JSON(http.StatusCreated, createUserRequest)
}

func NewUsersController(userService services.UsersService) UsersController {
	return &usersControllerImpl{
		userService,
	}
}
