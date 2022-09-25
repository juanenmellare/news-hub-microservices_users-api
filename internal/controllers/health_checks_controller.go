package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthChecksController interface {
	Ping(context *gin.Context)
}

type healthChecksController struct{}

func NewHealthChecksController() HealthChecksController {
	return &healthChecksController{}
}

func (h healthChecksController) Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
