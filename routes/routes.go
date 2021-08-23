package routes

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	kimiGroup := r.Group("/kimi")
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

	return r
}
