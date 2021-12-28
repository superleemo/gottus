package middlewares

import "github.com/gin-gonic/gin"

func ServerInitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Powered-By", "gottus")
		c.Next()
	}

}
