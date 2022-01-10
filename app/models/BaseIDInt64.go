package models

import (
	"fmt"
	"time"
)

type IBaseModelIDInt64 interface {
	Get() (interface{}, error)
}

type BaseModelIDInt64 struct {
	ID        int64      `gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:""`
}

func (b *BaseModelIDInt64) Get() (interface{}, error) {
	fmt.Println("base get")
	return nil, nil
}
