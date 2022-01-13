package captchamiddleware

import (
	"bytes"
	"fmt"
	captchacontroller "gin_graphql/app/controllers/CaptchaController"
	ginServices "gin_graphql/app/services/ginService"
	"gin_graphql/config/errorCode"
	"io/ioutil"

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
			ginServices.GinRespone(c, reqData, "", errorCode.FORBIDDEN, nil)
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
