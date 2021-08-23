package ginServices

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ResultObj struct {
	Status  bool
	Code    string
	Message string
}

type ResponseObj struct {
	Data   interface{}
	Result ResultObj
}

/**
 * @description:
 * @param {*gin.Context} c
 * @param {interface{}} reqJSON
 * @return {*}
 */
func GinRequest(c *gin.Context, reqJSON interface{}) (interface{}, error) {
	var reqData interface{}
	if c.Request.Method == "GET" {
		reqData = c.Request.URL.Query()
		if bindErr := c.ShouldBind(reqJSON); bindErr != nil {
			return nil, bindErr
		}
	} else {
		// 取得Body資料, 給log用
		data, err := c.GetRawData()
		if err != nil {
			return nil, err
		}
		reqData = ioutil.NopCloser(bytes.NewBuffer(data))

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		if bindErr := c.ShouldBind(reqJSON); bindErr != nil {
			return nil, bindErr
		}

	}
	return reqData, nil
}

/**
 * @description:
 * @param {*gin.Context} c
 * @param {interface{}} resquestData
 * @param {interface{}} responseData
 * @param {bool} resultStatus
 * @param {map[string]interface{}} resultMsg
 * @param {interface{}} err
 * @return {*}
 */
func GinRespone(c *gin.Context, resquestData interface{}, responseData interface{}, resultMsg map[string]interface{}, err interface{}) {
	var response ResponseObj
	if responseData == "" || responseData == nil {
		response.Data = nil
	} else {
		response.Data = responseData
	}
	env := os.Getenv("APP_ENV")
	response.Result = ResultObj{Status: resultMsg["status"].(bool), Code: resultMsg["code"].(string), Message: resultMsg["message"].(string)}

	if resultMsg["status"].(bool) {
		WriteLog(c, resquestData, response)
	} else {
		WriteLogError(c, resquestData, response, err)
	}

	if env == "app" {
		response.Result.Message = ""
	}

	c.JSON(resultMsg["httpCode"].(int), response)
}

/**
 * @description: 寫入 log
 * @param {*gin.Context} c
 * @param {interface{}} body
 * @param {interface{}} respon
 * @return {*}
 */
func WriteLog(c *gin.Context, body interface{}, respon interface{}) {
	log.WithFields(log.Fields{
		"url":        c.Request.URL.Path,
		"method":     c.Request.Method,
		"request_ip": c.ClientIP(),
		"body":       fmt.Sprintf("%+v", body),
	}).Info(fmt.Sprintf("%+v", respon))
}

/**
 * @description: 寫入 error log
 * @param {*gin.Context} c
 * @param {interface{}} body
 * @param {interface{}} respon
 * @param {interface{}} err
 * @return {*}
 */
func WriteLogError(c *gin.Context, body interface{}, respon interface{}, err interface{}) {
	log.WithFields(log.Fields{
		"url":        c.Request.URL.Path,
		"method":     c.Request.Method,
		"request_ip": c.ClientIP(),
		"body":       fmt.Sprintf("%+v", body),
		"err_msg":    err,
	}).Error(fmt.Sprintf("%+v", respon))
}
