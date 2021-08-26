package models

import (
	"gin_graphql/config/mysql"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = mysql.GormDB
}

type BaseModel struct {
	ID        int        `json:"id" form:"id,omitempty" structs:"id,omitempty" gorm:"type:int;primary_key"`
	CreatedAt time.Time  `gorm:"index;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"index"`
}
