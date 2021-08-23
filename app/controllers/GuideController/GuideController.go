package GuideController

import (
	"fmt"
	ginServices "gin_restful_graphql/app/services/ginService"
	"gin_restful_graphql/config/errorCode"

	"github.com/gin-gonic/gin"
)

/**
 * @description: GET
 * @param {*gin.Context} c
 * @return {*}
 */
func GetGuide(c *gin.Context) {
	// 定義接收格式
	type structRequest struct{}
	var reqJSON structRequest
	reqData, err := ginServices.GinRequest(c, &reqJSON)
	if err != nil {
		ginServices.GinRespone(c, reqData, "", errorCode.PARAMS_INVALID, err)
		return
	}
	// TODO: Query DB
	id := c.Param("id")
	fmt.Println(id)
	result := map[string]interface{}{
		"id": id,
	}
	ginServices.GinRespone(c, reqData, result, errorCode.SUCCESS, "")
}

/**
 * @description: POST
 * @param {*gin.Context} c
 * @return {*}
 */
func PostGuide(c *gin.Context) {
	// 定義接收格式
	type structRequest struct{}
	var reqJSON structRequest
	reqData, err := ginServices.GinRequest(c, &reqJSON)
	if err != nil {
		ginServices.GinRespone(c, reqData, "", errorCode.PARAMS_INVALID, err)
		return
	}
	// TODO: Create

	result := map[string]interface{}{
		"id": 5566,
	}
	ginServices.GinRespone(c, reqData, result, errorCode.SUCCESS, "")
}

/**
 * @description: PUT
 * @param {*gin.Context} c
 * @return {*}
 */
func PutGuide(c *gin.Context) {
	// 定義接收格式
	type structRequest struct{}
	var reqJSON structRequest
	reqData, err := ginServices.GinRequest(c, &reqJSON)
	if err != nil {
		ginServices.GinRespone(c, reqData, "", errorCode.PARAMS_INVALID, err)
		return
	}
	// TODO: Update
	id := c.Param("id")
	fmt.Println(id)
	result := map[string]interface{}{
		"id": id,
	}
	ginServices.GinRespone(c, reqData, result, errorCode.SUCCESS, "")
}

/**
 * @description: DELETE
 * @param {*gin.Context} c
 * @return {*}
 */
func DeleteGuide(c *gin.Context) {
	// 定義接收格式
	type structRequest struct{}
	var reqJSON structRequest
	reqData, err := ginServices.GinRequest(c, &reqJSON)
	if err != nil {
		ginServices.GinRespone(c, reqData, "", errorCode.PARAMS_INVALID, err)
		return
	}
	// TODO: Delete
	id := c.Param("id")
	fmt.Println(id)
	result := map[string]interface{}{
		"id": id,
	}
	ginServices.GinRespone(c, reqData, result, errorCode.SUCCESS, "")
}
