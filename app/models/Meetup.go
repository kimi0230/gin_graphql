package models

import (
	"fmt"
)

type Meetup struct {
	BaseModel
	Name        string `json:"name" form:"name,omitempty" structs:"name,omitempty" gorm:"Column:name;type:varchar(32);comment:'name' "`
	Description string `json:"description" form:"description,omitempty" structs:"description,omitempty" gorm:"Column:description;type:varchar(255);comment:'description' "`
	UserID      int    `json:"userId" form:"userId,omitempty" structs:"userId,omitempty"`
}

func (Meetup) TableName() string {
	return "meetup"
}

func (m *Meetup) Create(meetup *Meetup) (*Meetup, error) {
	if err := DB.Create(&meetup).Error; err != nil {
		return nil, err
	}
	return meetup, nil
}

func (m *Meetup) Update(id int, input interface{}) (*Meetup, error) {
	tx := DB.Begin()
	query := tx.Model(m).Where("id = ?", id).Updates(input)
	if query.Error != nil {
		tx.Rollback()
		return nil, query.Error
	}
	tx.Commit()
	return m, nil
}

func (m *Meetup) Delete(id int) (bool, error) {
	tx := DB.Begin()
	fmt.Println("--->", id)
	if err := tx.Where("id = ?", id).Delete(m).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}

func (m *Meetup) GetMeetupsByUser(user *User) ([]*Meetup, error) {
	var meetups []*Meetup
	if err := DB.Model(&meetups).Where("user_id = ?", user.ID).Find(&meetups).Order("id").Error; err != nil {
		return nil, err
	}
	return meetups, nil
}
