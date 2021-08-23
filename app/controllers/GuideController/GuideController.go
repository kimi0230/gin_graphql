package GuideController

import (
	"fmt"
	services "gin_restful_graphql/app/services/ginService"
	"gin_restful_graphql/config/errorCode"

	"github.com/gin-gonic/gin"
)

func GetGuide(c *gin.Context) {
	// 定義接收格式
	type structRequest struct{}
	var reqJSON structRequest
	reqData, err := services.GinRequest(c, &reqJSON)
	if err != nil {
		services.GinRespone(c, reqData, "", false, errorCode.PARAMS_INVALID, err)
		return
	}
	// TODO: Query DB
	id := c.Param("id")
	fmt.Println(id)
	result := map[string]interface{}{
		"id": id,
	}
	services.GinRespone(c, reqData, result, true, errorCode.SUCCESS, "")
}
