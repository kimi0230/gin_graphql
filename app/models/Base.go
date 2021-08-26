package models

import (
	"fmt"
	"gin_graphql/config/mysql"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = mysql.GormDB
}

type IBaseModel interface {
	Get() (interface{}, error)
}

type BaseModel struct {
	ID        int        `json:"id" form:"id,omitempty" structs:"id,omitempty"`
	CreatedAt time.Time  `gorm:"index;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"index"`
}

func (b *BaseModel) Get() (interface{}, error) {
	fmt.Println("base get")
	return nil, nil
}
