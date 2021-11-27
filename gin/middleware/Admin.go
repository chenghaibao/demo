package middleware

import (
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		success := "success"
		if success == "success" {
			c.Next()
		}
		c.Abort()
	}
}
