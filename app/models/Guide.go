package models

import "time"

type Guide struct {
	ID        int       `json:"id" form:"id,omitempty" structs:"id,omitempty"`
	Code      string    `json:"code" form:"code,omitempty" structs:"code,omitempty" gorm:"Column:code;type:varchar(32);comment:'code' "`
	Project   string    `json:"project" form:"project,omitempty" structs:"project,omitempty" gorm:"Column:project;type:varchar(32);comment:'project 代碼' "`
	ZH        string    `json:"zh" form:"zh,omitempty" structs:"zh,omitempty" gorm:"Column:zh;default:'';type:text;comment:'' "`
	EN        string    `json:"en" form:"en,omitempty" structs:"en,omitempty" gorm:"Column:en;default:'';type:text;comment:'' "`
	JP        string    `json:"jp" form:"jp,omitempty" structs:"jp,omitempty" gorm:"Column:jp;default:'';type:text;comment:'' "`
	TH        string    `json:"th" form:"th,omitempty" structs:"th,omitempty" gorm:"Column:th;default:'';type:text;comment:'' "`
	JA        string    `json:"ja" form:"ja,omitempty" structs:"ja,omitempty" gorm:"Column:ja;default:'';type:text;comment:'' "`
	DE        string    `json:"de" form:"de,omitempty" structs:"de,omitempty" gorm:"Column:de;default:'';type:text;comment:'' "`
	FR        string    `json:"fr" form:"fr,omitempty" structs:"fr,omitempty" gorm:"Column:fr;default:'';type:text;comment:'' "`
	It        string    `json:"it" form:"it,omitempty" structs:"it,omitempty" gorm:"Column:it;default:'';type:text;comment:'' "`
	Pt        string    `json:"pt" form:"pt,omitempty" structs:"pt,omitempty" gorm:"Column:pt;default:'';type:text;comment:'' "`
	Es        string    `json:"es" form:"es,omitempty" structs:"es,omitempty" gorm:"Column:es;default:'';type:text;comment:'' "`
	CreatedAt time.Time `json:"createdAt" gorm:"Column:createdAt;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"Column:updatedAt;default:CURRENT_TIMESTAMP"`
}

// 自訂對應的table name
func (Guide) TableName() string {
	return "guide"
}
