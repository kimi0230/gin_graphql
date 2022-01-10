package models

type StaffRole struct {
	BaseModelIDInt64
	StaffID int64 `json:"staff_id" gorm:"Column:staff_id;type:bigint;not null;comment:'staff id'" `
	RoleID  int   `json:"role_id" gorm:"Column:role_id;type:int;not null;comment:'role id'" `
}

func (StaffRole) TableName() string {
	return "staff_role"
}
