package server

import (
	"github.com/gin-gonic/gin"
	"github.com/superleemo/gottus/middlewares"
)

func Start() {
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong"})
	})
	server.Use(middlewares.ServerInitMiddleware())
	server.Run()
}

func ServerInitMiddleware() {
	panic("unimplemented")
}
