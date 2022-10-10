package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"news-hub-microservices_users-api/internal/errors"
	"news-hub-microservices_users-api/internal/rest"
	"news-hub-microservices_users-api/internal/services"
)

type UsersController interface {
	Create(context *gin.Context)
	Authenticate(context *gin.Context)
	Get(context *gin.Context)
}

type usersController struct {
	userService              services.UsersService
	userTokenSecretKey       string
	userTokenExpirationHours int
}

func (u usersController) Create(context *gin.Context) {
	var request rest.CreateUserRequest
	request.MarshallAndValidate(context)

	uuid := u.userService.Create(*request.FirstName, *request.LastName, *request.Email, *request.Password)

	fmt.Println("5")

	response := rest.NewCreateUserResponse(uuid)
	fmt.Println("6")

	context.JSON(http.StatusCreated, response)
	context.Done()
}

func (u usersController) Authenticate(context *gin.Context) {
	var request rest.AuthenticateRequest
	request.MarshallAndValidate(context)

	user := u.userService.Authenticate(*request.Email, *request.Password)
	fmt.Println("5")

	token := rest.NewUserToken(8, user)
	fmt.Println("6")

	response := rest.NewAuthenticateResponse(token.ToString(u.userTokenSecretKey))

	fmt.Println("7")

	context.JSON(http.StatusOK, &response)
	context.Done()
}

func (u usersController) Get(context *gin.Context) {
	userToken := &rest.UserToken{}
	userToken.Verify(u.userTokenSecretKey, context.Request)
	userToken.IsExpired()

	user := u.userService.GetById(userToken.GetUserId())
	if user == nil {
		panic(errors.NewNotFoundError("user from token not found"))
	}

	response := rest.NewGetResponse(user)

	context.JSON(http.StatusOK, &response)
	context.Done()
}

func NewUsersController(userService services.UsersService, userTokenSecretKey string, userTokenExpirationHours int) UsersController {
	return &usersController{
		userService,
		userTokenSecretKey,
		userTokenExpirationHours,
	}
}
