package server

import "github.com/gin-gonic/gin"

func Start() {
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong"})
	})
	server.Run()
}
