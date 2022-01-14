package routes

import (
	captchacontroller "gin_graphql/app/controllers/CaptchaController"
	"gin_graphql/app/controllers/GuideController"
	"gin_graphql/app/controllers/LoginoutController"
	"gin_graphql/app/controllers/MeetupController"
	"gin_graphql/app/middleware/captchamiddleware"
	"gin_graphql/app/middleware/headerAuth"
	"gin_graphql/app/middleware/rateLimit"
	"gin_graphql/app/middleware/staffRoleAuth"
	jwtservice "gin_graphql/app/services/jwtService"

	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api", rateLimit.RateLimitToken())
	apiGroup.POST("/login", LoginoutController.Login)
	apiGroup.POST("/login/test", jwtservice.VerifyJWTAuth(), LoginoutController.Test)

	// RESTful >>>
	// v1 >>>
	apiv1Group := apiGroup.Group("/v1", headerAuth.VerifyHeaderAuth())
	apiv1Group.GET("/guide", GuideController.GetGuide)
	apiv1Group.GET("/guide/:id", GuideController.GetGuide)
	apiv1Group.POST("/guide", GuideController.PostGuide)
	apiv1Group.PUT("/guide/:id", GuideController.PutGuide)
	apiv1Group.DELETE("/guide/:id", GuideController.DeleteGuide)
	apiv1Group.GET("/meetups", MeetupController.GetMeetup)
	apiv1Group.GET("/meetups/:id", MeetupController.GetMeetup)
	// v1 <<<
	// RESTful <<<

	// GraphQL >>>

	// GraphQL <<<

	// 驗證碼 >>>
	// dchest/captcha
	captchaGroup := r.Group("/captcha")
	// 取得驗證碼圖片路徑
	captchaGroup.GET("/", captchacontroller.NewCaptchaString)
	// 取得驗證碼圖片檔案
	captchaGroup.GET("/:captchaId", captchacontroller.GetCaptcha)
	// demo 驗證
	captchaGroup.POST("/verify", captchamiddleware.Verify(), captchacontroller.VerifyCaptcha)
	// 驗證碼 <<<

	// Kimi 測試區 >>>
	kimiGroup := r.Group("/kimi", rateLimit.RateLimitLeaky())
	kimiGroup.GET("/role", staffRoleAuth.VerifyStaffAuth(), func(c *gin.Context) {
		// 測試 staff role 驗證 中間層
		c.String(http.StatusOK, "ok")
	})
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
