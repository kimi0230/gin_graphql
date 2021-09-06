package middleware

import (
	"context"
	"errors"
	"fmt"
	"gin_graphql/app/models"
	ginServices "gin_graphql/app/services/ginService"
	"gin_graphql/config/errorCode"
	"strings"

	"github.com/gin-gonic/gin"
)

const CurrentUserKey = "currentUser"

type HeaderValid struct {
	Authorization string `json:"authorization" form:"authorization"  binding:"required"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// TODO: 檢查 Header
		var reqJSON HeaderValid
		// if bindErr := c.ShouldBindHeader(&reqJSON); bindErr != nil {
		// 	c.Abort()
		// 	return
		// }

		token, err := stripBearerPrefixFromToken(reqJSON.Authorization)
		if err != nil {
			ginServices.GinRespone(c, "", "", errorCode.FORBIDDEN, nil)
			c.Abort()
			return
		}

		// TODO: 檢查token
		fmt.Println("token ", token)

		// TODO: 取得 User
		user := &models.User{}
		user, _ = user.GetUserByID(2)
		// c.Set(CurrentUserKey, user)

		// 寫入 graphql 的 context. 再包進gin
		ctx := context.WithValue(c.Request.Context(), CurrentUserKey, user)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func stripBearerPrefixFromToken(token string) (string, error) {
	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, nil
}

func GetCurrentUserFromCTX(ctx context.Context) (*models.User, error) {
	errNoUserInContext := errors.New("no user in context")
	fmt.Println("ctx.Value(CurrentUserKey) ", ctx.Value(CurrentUserKey))
	if ctx.Value(CurrentUserKey) == nil {
		return nil, errNoUserInContext
	}

	user, ok := ctx.Value(CurrentUserKey).(*models.User)

	if !ok || user.ID == 0 {
		return nil, errNoUserInContext
	}

	return user, nil
}
