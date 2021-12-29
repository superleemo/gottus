package server

import (
	"github.com/gin-gonic/gin"
	"github.com/superleemo/gottus/controllers"
	"github.com/superleemo/gottus/middlewares"
)

func Start() {

	health := &controllers.HealthController{}
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong"})
	})
	server.GET("/health", health.Status)
	server.Use(middlewares.ServerInitMiddleware())
	server.Run()
}

func ServerInitMiddleware() {
	panic("unimplemented")
}
