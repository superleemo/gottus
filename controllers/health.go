package controllers

import (
	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	// c.String(http.StatusOK, "Working!")
	c.JSON(200, gin.H{
		"status": "success",
		"body":   "working",
	})
}
