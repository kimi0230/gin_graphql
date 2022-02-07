package staffRoleAuth

import (
	"fmt"
	ginServices "gin_graphql/app/services/ginService"
	"gin_graphql/config/errorCode"
	"gin_graphql/config/mysql"

	"github.com/gin-gonic/gin"
)

func VerifyStaffAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 測試, 使用者的 role_id
		c.Set("roleID", "1")

		roleID, err := c.Get("roleID")
		if !err {
			ginServices.GinRespone(c, "", "", errorCode.BAD_REQUEST, nil)
			c.Abort()
			return
		}

		// 檢查 role id 在 permission 權限
		// SELECT rp.role_id, rp.permission_id, p.title, p.slug, p.active
		// FROM role_permission as rp
		// JOIN permission as p ON p.id=rp.permission_id
		// WHERE p.active="1";
		type permissionStruct struct {
			RoleID       int    `json:"role_id" form:"role_id,omitempty" structs:"role_id,omitempty" `
			PermissionID int    `json:"permission_id" form:"permission_id,omitempty"`
			Title        string `json:"title" form:"title,omitempty" `
			Slug         string `json:"slug" form:"slug,omitempty" structs:"slug,omitempty" `
			Active       string `json:"active" form:"active,omitempty"`
		}
		var permission permissionStruct
		sel := `rp.role_id, rp.permission_id, p.title, p.slug, p.active`
		if err := mysql.GormDB.Table("role_permission as rp").Select(sel).Joins("join permission as p ON p.id=rp.permission_id").Where("p.active=? and p.slug=?", "1", c.Request.URL.Path).Scan(&permission).Error; err != nil {
			ginServices.GinRespone(c, "", "", errorCode.FORBIDDEN, nil)
			c.Abort()
			return
		}
		fmt.Println("roleID ===>", roleID)
		fmt.Println("URL.Path ====>", c.Request.URL.Path)
		c.Next()
	}
}
