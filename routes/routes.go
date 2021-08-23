package routes

import (
	"gin_restful_graphql/app/controllers/GuideController"
	"gin_restful_graphql/app/middleware/headerAuth"
	"gin_restful_graphql/app/middleware/rateLimit"

	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// RESTful >>>
	// v1 >>>
	apiv1Group := r.Group("/api/v1", rateLimit.RateLimitToken(), headerAuth.VerifyHeaderAuth())
	apiv1Group.GET("/guide", GuideController.GetGuide)
	apiv1Group.GET("/guide/:id", GuideController.GetGuide)
	apiv1Group.POST("/guide", GuideController.PostGuide)
	apiv1Group.PUT("/guide/:id", GuideController.PutGuide)
	apiv1Group.DELETE("/guide/:id", GuideController.DeleteGuide)
	// v1 <<<
	// RESTful <<<

	// GraphQL >>>

	// GraphQL <<<

	// Kimi 測試區 >>>
	kimiGroup := r.Group("/kimi", rateLimit.RateLimitLeaky())
	kimiGroup.GET("/", func(c *gin.Context) {
		result := map[string]interface{}{
			"Data": map[string]interface{}{
				"Msg":     "Hello Kimi!",
				"AppName": os.Getenv("APP_NAME"),
				"AppENV":  os.Getenv("APP_ENV"),
				"Time":    time.Now().Format("2006-01-02 15:04:05"),
				"TimeUTC": time.Now().UTC().Format("2006-01-02 15:04:05"),
			},
			"Result": map[string]interface{}{
				"Status":  true,
				"Code":    "2000",
				"Message": "正確狀況",
			},
		}
		c.JSON(http.StatusOK, result)
	})
	// Kimi 測試區 <<<

	return r
}
