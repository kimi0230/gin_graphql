package models

type Permission struct {
	BaseModel
	Title       string `json:"title" gorm:"Column:title;type:varchar(50);not null;comment:'權限名稱'" `
	Slug        string `json:"slug" gorm:"Column:slug;type:varchar(100);not null;comment:'slug'" `
	Description string `json:"description" gorm:"Column:description;type:tinytext;not null;comment:'權限描述'" `
	Active      string `json:"active" gorm:"Column:active;type:varchar(1);not null;default:'1';comment:'1:work,0:no work'" `
}

func (Permission) TableName() string {
	return "permission"
}
