package rateLimit

import (
	ginServices "gin_restful_graphql/app/services/ginService"
	"gin_restful_graphql/config/errorCode"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

var leakyLimit = ratelimit.New(100000, ratelimit.WithoutSlack) // per second
func RateLimitLeaky() gin.HandlerFunc {
	return func(c *gin.Context) {
		for i := 0; i < 1; i++ {
			leakyLimit.Take()
		}
		c.Next()
	}
}

var tokenLimit = tollbooth.NewLimiter(100000, nil) // per second

func RateLimitToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// tokenLimit.SetMessage("You have reached maximum request limit.")
		httpError := tollbooth.LimitByRequest(tokenLimit, c.Writer, c.Request)
		if httpError != nil {
			ginServices.GinRespone(c, "", "", errorCode.BAD_REQUEST, httpError)
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
