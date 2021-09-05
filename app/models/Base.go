package models

import (
	"fmt"
	"gin_graphql/config/mysql"
	"time"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	DB = mysql.GormDB
}

type IBaseModel interface {
	Get() (interface{}, error)
}

type BaseModel struct {
	ID        int        `json:"id" form:"id,omitempty" structs:"id,omitempty"`
	CreatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:""`
}

func (b *BaseModel) Get() (interface{}, error) {
	fmt.Println("base get")
	return nil, nil
}
