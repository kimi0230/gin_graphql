package models

import (
	"fmt"
	"gin_graphql/graph/model"
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

func (m *Meetup) Get(filter *model.MeetupFilter, limit, offset *int) ([]*Meetup, error) {
	var meetups []*Meetup
	query := db.Model(&meetups).Order("id DESC")
	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
		}
	}

	if limit != nil {
		query = query.Limit(*limit)
	}

	if offset != nil {
		query = query.Offset(*offset)
	}

	if err := query.Find(&meetups).Error; err != nil {
		return nil, err
	}

	return meetups, nil
}

func (m *Meetup) Create(meetup *Meetup) (*Meetup, error) {
	if err := db.Create(&meetup).Error; err != nil {
		return nil, err
	}
	return meetup, nil
}

func (m *Meetup) Update(id int, input interface{}) (*Meetup, error) {
	tx := db.Begin()
	query := tx.Model(m).Where("id = ?", id).Updates(input)
	if query.Error != nil {
		tx.Rollback()
		return nil, query.Error
	}
	tx.Commit()
	return m, nil
}

func (m *Meetup) Delete(id int) (bool, error) {
	tx := db.Begin()
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
	if err := db.Model(&meetups).Where("user_id = ?", user.ID).Find(&meetups).Order("id").Error; err != nil {
		return nil, err
	}
	return meetups, nil
}
