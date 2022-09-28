package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"news-hub-microservices_users-api/configs"
	"news-hub-microservices_users-api/internal/errors"
	"news-hub-microservices_users-api/internal/factories"
)

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

func BasicAuthenticationMiddleware(c *gin.Context, candidateUsername, candidatePassword string) {
	username, password, hasAuth := c.Request.BasicAuth()
	if !hasAuth || username != candidateUsername || password != candidatePassword {
		c.JSON(http.StatusUnauthorized, errors.NewUnauthorizedApiError("invalid credentials"))
		c.Abort()
	}
}

func NewRouter(controllers factories.ControllersFactory, config configs.Config) *gin.Engine {
	router := gin.Default()
	router.Use(gin.CustomRecovery(HandlePanicRecoveryMiddleware))

	router.GET("/ping", controllers.GetHealthChecksController().Ping)

	usersController := controllers.GetUsersController()
	v1 := router.Group("/v1")
	v1.Use(func(c *gin.Context) {
		BasicAuthenticationMiddleware(c, config.GetBasicAuthUsername(), config.GetBasicAuthPassword())
	})
	{
		v1.GET("", usersController.Get)
		v1.POST("", usersController.Create)
		v1.POST("/login", usersController.Authenticate)
	}

	return router
}
