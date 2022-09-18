package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news-hub-microservices_users-api/api"
	"news-hub-microservices_users-api/services"
)

type UsersController interface {
	Create(context *gin.Context)
	Authenticate(context *gin.Context)
}

type usersControllerImpl struct {
	userService services.UsersService
}

func (u usersControllerImpl) Create(context *gin.Context) {
	var request api.CreateUserRequest
	request.MarshallAndValidate(context)

	uuid := u.userService.Create(*request.FirstName, *request.LastName, *request.Email, *request.Password)

	response := api.NewCreateUserResponse(uuid)

	context.JSON(http.StatusCreated, response)
}

func (u usersControllerImpl) Authenticate(context *gin.Context) {
	var request api.AuthenticateRequest
	request.MarshallAndValidate(context)

	user := u.userService.Authenticate(*request.Email, *request.Password)

	token := api.NewUserToken(8, user)

	response := api.NewAuthenticateResponse(token, "hello_world")

	context.JSON(http.StatusOK, &response)
}

func NewUsersController(userService services.UsersService) UsersController {
	return &usersControllerImpl{
		userService,
	}
}
