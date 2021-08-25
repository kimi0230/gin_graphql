package models

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Meetups   []*Meetup `json:"meetups"`
}
