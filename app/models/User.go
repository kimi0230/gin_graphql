package models

type User struct {
	BaseModel
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Meetups   []*Meetup `json:"meetups"`
}
