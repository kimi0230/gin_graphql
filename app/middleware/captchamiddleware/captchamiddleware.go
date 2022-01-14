package captchamiddleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	captchacontroller "gin_graphql/app/controllers/CaptchaController"
	curlservice "gin_graphql/app/services/curlService"
	ginServices "gin_graphql/app/services/ginService"
	"gin_graphql/config/errorCode"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqJSON captchacontroller.CaptchaRequest
		var reqData interface{}

		// 取得body資料
		data, err := c.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}
		reqData = ioutil.NopCloser(bytes.NewBuffer(data))
		// 因為body只能取一次,再把request補回去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		if bindErr := c.ShouldBind(&reqJSON); bindErr != nil {
			ginServices.GinRespone(c, reqData, "", errorCode.PARAMS_INVALID, nil)
			c.Abort()
			return
		}
		// 因為body只能取一次,再把request補回去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		// 驗證captcha
		// The function deletes the captcha with the given id from the internal storage, so that the same captcha can't be verified anymore.
		if !captchacontroller.VerifyString(reqJSON) {
			ginServices.GinRespone(c, reqData, "", errorCode.FORBIDDEN, nil)
			c.Abort()
			return

		}

		c.Next()
	}
}

func VerifyreCAPTCHA() gin.HandlerFunc {
	return func(c *gin.Context) {
		type reCaptchaReq struct {
			Grecaptcha string `json:"g_recaptcha" form:"g_recaptcha"  binding:"required"`
		}
		type reCaptchaRes struct {
			Success    bool     `json:"success" form:"success"  binding:"required"`
			ErrorCodes []string `json:"error-codes" form:"error-codes"  binding:"required"`
		}

		var reqJSON reCaptchaReq
		var reqData interface{}

		// 取得body資料
		data, err := c.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}
		reqData = ioutil.NopCloser(bytes.NewBuffer(data))
		// 因為body只能取一次,再把request補回去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		if bindErr := c.ShouldBind(&reqJSON); bindErr != nil {
			ginServices.GinRespone(c, reqData, "", errorCode.PARAMS_INVALID, nil)
			c.Abort()
			return
		}
		// 因為body只能取一次,再把request補回去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		// 驗證 google reCcaptcha
		url := "https://www.google.com/recaptcha/api/siteverify"
		sendJSON := map[string]interface{}{
			"secret":   os.Getenv("RECAPTCHA_SECRET"),
			"response": reqJSON.Grecaptcha,
		}
		respBody, ok := curlservice.PostJSON(url, sendJSON)
		if !ok {
			ginServices.GinRespone(c, reqData, "", errorCode.PARAMS_INVALID, err)
			c.Abort()
			return
		}
		// 轉 struct
		var resultJSON reCaptchaRes
		if err := json.Unmarshal([]byte(respBody), &resultJSON); err != nil {
			ginServices.GinRespone(c, reqData, "", errorCode.PARAMS_INVALID, err)
			c.Abort()
			return
		}

		if !resultJSON.Success {
			ginServices.GinRespone(c, reqData, "", errorCode.PARAMS_INVALID, resultJSON.ErrorCodes)
			c.Abort()
			return
		}

		c.Next()
	}

}
