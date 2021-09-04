package middleware

import (
	"fmt"
	ginServices "gin_graphql/app/services/ginService"
	"gin_graphql/config/errorCode"
	"strings"

	"github.com/gin-gonic/gin"
)

type HeaderValid struct {
	Authorization string `json:"authorization" form:"authorization"  binding:"required"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqJSON HeaderValid
		// 檢查 Header
		if bindErr := c.ShouldBindHeader(&reqJSON); bindErr != nil {
			c.Abort()
			return
		}

		headerAuth := strings.Split(reqJSON.Authorization, " ")
		if headerAuth[0] != "Bearer" {
			ginServices.GinRespone(c, "", "", errorCode.FORBIDDEN, nil)
			c.Abort()
			return
		}
		token := headerAuth[1]
		fmt.Println("token ", token)
		// TODO: 檢查token

		c.Next()
	}
}
