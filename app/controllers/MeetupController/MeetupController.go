package MeetupController

import (
	"fmt"
	ginServices "gin_graphql/app/services/ginService"
	"gin_graphql/config/errorCode"

	"github.com/gin-gonic/gin"
)

// GetMeetup godoc
// @Summary Show an meetup
// @Description get string by ID
// @Tags meetups
// @Accept  json
// @Produce  json
// @param Authorization header string true "Authorization"
// @Param id path int true "Meetup ID"
// @Header 200 {string} Token "qwerty"
// @Success 200 {object} ginServices.ResponseObj{data=[]models.Meetup} "desc"
// @Router /meetups/{id} [get]
func GetMeetup(c *gin.Context) {
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
