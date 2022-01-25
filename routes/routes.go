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
	"gin_graphql/config/corsConfig"
	"gin_graphql/routes/kimi"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(cors.Default())
	r.Use(cors.New(corsConfig.CorsConfig()))

	/***    api v1   ***/
	apiGroup := r.Group("/api", rateLimit.RateLimitToken())
	apiGroup.POST("/login", LoginoutController.Login)
	apiGroup.POST("/login/test", jwtservice.VerifyJWTAuth(), LoginoutController.Test)

	/***   RESTful  ***/
	apiv1Group := apiGroup.Group("/v1", headerAuth.VerifyHeaderAuth())
	apiv1Group.GET("/guide", GuideController.GetGuide)
	apiv1Group.GET("/guide/:id", GuideController.GetGuide)
	apiv1Group.POST("/guide", GuideController.PostGuide)
	apiv1Group.PUT("/guide/:id", GuideController.PutGuide)
	apiv1Group.DELETE("/guide/:id", GuideController.DeleteGuide)
	apiv1Group.GET("/meetups", MeetupController.GetMeetup)
	apiv1Group.GET("/meetups/:id", MeetupController.GetMeetup)
	/*******************/

	/***   TODO:GraphQL  ***/

	/*******************/

	/***   Captcha  ***/
	// dchest/captcha
	captchaGroup := r.Group("/captcha")
	// 取得驗證碼圖片路徑
	captchaGroup.GET("/", captchacontroller.NewCaptchaString)
	// 取得驗證碼圖片檔案
	captchaGroup.GET("/:captchaId", captchacontroller.GetCaptcha)
	// demo 驗證
	captchaGroup.POST("/verify", captchamiddleware.Verify(), captchacontroller.VerifyCaptcha)
	/*******************/

	/***  Kimi 測試區 ***/
	kimiGroup := r.Group("/kimi", rateLimit.RateLimitLeaky())
	// kimiGroup.Use(cors.New(corsConfig.CorsConfig()))
	kimiGroup.GET("/", kimi.Info)
	kimiGroup.GET("/role", staffRoleAuth.VerifyStaffAuth(), kimi.Ok)
	/*******************/

	return r
}
