package middleware

import (
	"fmt"
	ginServices "gin_restful_graphql/app/services/ginService"
	"gin_restful_graphql/config/errorCode"
	"strings"

	"github.com/gin-gonic/gin"
)

type HeaderValid struct {
	Authorization string `json:"authorization" form:"authorization"  binding:"required"`
	Lang          string `json:"lang" form:"lang" binding:"-"`
	UserZone      string `json:"user_zone" form:"user_zone" binding:"-"`
}

func VerifyHeaderAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqData interface{}
		var reqJSON HeaderValid

		// 檢查 Header
		if bindErr := c.ShouldBindHeader(&reqJSON); bindErr != nil {
			ginServices.GinRespone(c, reqData, "", errorCode.PARAMS_INVALID, bindErr)
			c.Abort()
			return
		}
		headerAuth := strings.Split(reqJSON.Authorization, " ")
		if headerAuth[0] != "Bearer" {
			ginServices.GinRespone(c, reqData, "", errorCode.FORBIDDEN, nil)
			c.Abort()
			return
		}
		token := headerAuth[1]

		// TODO: 檢查token
		fmt.Println("token =  ", token)
		fmt.Println("Lang =  ", reqJSON.Lang)
		/*
			// 取出 body
			if c.Request.Method == "GET" {
				reqData = c.Request.URL.Query()
			} else {
				// 取得body資料
				data, err := c.GetRawData()
				if err != nil {
					fmt.Println(err.Error())
				}
				reqData = ioutil.NopCloser(bytes.NewBuffer(data))
				// 因為body只能取一次,再把request補回去
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

				if bindErr := c.ShouldBind(&reqJSON); bindErr != nil {
					ginServices.GinRespone(c, reqData, "", errorCode.PARAMS_INVALID, bindErr)
					c.Abort()
					return
				}
				// 因為body只能取一次,再把request補回去
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			}
		*/

		c.Set("token", token)
		c.Next()
	}
}
