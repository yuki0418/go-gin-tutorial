package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yuki0418/go-gin-tutorial/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	healthController := new(controllers.HealthController)

	router.GET("/health", healthController.Status)

	return router
}