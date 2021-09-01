package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("kkk")
		// c.Abort()
		c.Next()
	}
}
