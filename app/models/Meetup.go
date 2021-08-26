package models

type Meetup struct {
	BaseModel
	Name        string `json:"name" form:"name,omitempty" structs:"name,omitempty" gorm:"Column:name;type:varchar(32);comment:'name' "`
	Description string `json:"description" form:"description,omitempty" structs:"description,omitempty" gorm:"Column:description;type:varchar(255);comment:'description' "`
	UserID      int    `json:"userId" form:"userId,omitempty" structs:"userId,omitempty"`
}

func (Meetup) TableName() string {
	return "meetup"
}

func (m *Meetup) Get() ([]*Meetup, error) {
	var meetups []*Meetup
	if err := db.Model(meetups).Find(&meetups).Error; err != nil {
		return nil, err
	}
	return meetups, nil
}
