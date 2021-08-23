package GuideController

import (
	"fmt"
	ginServices "gin_restful_graphql/app/services/ginService"
	"gin_restful_graphql/config/errorCode"

	"github.com/gin-gonic/gin"
)

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
