package router

import "news-hub-microservices_users-api/factories"
import "github.com/gin-gonic/gin"

func New(controllers factories.ControllersFactory) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.GetHealthChecksController().Ping)

	return router
}
