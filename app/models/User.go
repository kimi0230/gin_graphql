package models

import "fmt"

type User struct {
	BaseModel
	Account  string `json:"account"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserByField(field, value string) (*User, error) {
	var user *User
	if err := DB.Model(user).Where(fmt.Sprintf("%v = ?", field), value).First(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *User) GetUserByID(id int) (*User, error) {
	user := &User{}
	if err := DB.Model(user).Where("id = ?", id).First(user).Error; err != nil {
		fmt.Println("kkkkk")
		return user, err
	}
	return user, nil
}

func (u *User) Get() (*User, error) {
	users := &User{}
	if err := DB.Model(users).First(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
