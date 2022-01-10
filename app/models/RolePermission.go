package models

// 將 role 的權限新增刪除都用這張表
type RolePermission struct {
	BaseModel
	RoleID       int   `json:"role_id" gorm:"Column:role_id;type:int;not null;comment:'role id'" `
	PermissionID int64 `json:"permission_id" gorm:"Column:permission_id;type:int;not null;comment:'permission_id'" `
}

func (RolePermission) TableName() string {
	return "role_permission"
}
