package router

import (
	"fmt"
	"news-hub-microservices_users-api/errors"
	"news-hub-microservices_users-api/factories"
)
import "github.com/gin-gonic/gin"

func HandlePanicRecoveryMiddleware(context *gin.Context, i interface{}) {
	var apiError *errors.ApiError
	switch err := i.(type) {
	case *errors.ApiError:
		apiError = err
	case *errors.AlreadyExistModelError:
		apiError = errors.NewBadRequestApiError(err.Message)
	case *errors.InvalidEmailOrPasswordError:
		apiError = errors.NewBadRequestApiError(err.Message)
	case error:
		apiError = errors.NewInternalServerApiError(fmt.Sprintf("unexpected error: %v", err))
	default:
		apiError = errors.NewInternalServerApiError(fmt.Sprintf("unhandled error: %v", err))
	}
	context.JSON(apiError.Code, apiError)
}

func New(controllers factories.ControllersFactory) *gin.Engine {
	router := gin.Default()
	router.Use(gin.CustomRecovery(HandlePanicRecoveryMiddleware))

	router.GET("/ping", controllers.GetHealthChecksController().Ping)

	usersController := controllers.GetUsersController()
	router.POST("/", usersController.Create)
	router.POST("/login", usersController.Authenticate)
	router.GET("/", usersController.Get)

	return router
}
