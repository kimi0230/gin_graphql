package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Account  string `json:"account"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) HashPassword(password string) error {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(passwordHash)

	return nil
}

func (u *User) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (u *User) Create(user *User) (*User, error) {
	if err := DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) GetUserByField(field, value string) (*User, error) {
	user := &User{}
	if err := DB.Model(user).Where(fmt.Sprintf("%v = ?", field), value).Take(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *User) GetUserByID(id int) (*User, error) {
	user := &User{}
	if err := DB.Model(user).Where("id = ?", id).First(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *User) GetUserByEmail(email string) (*User, error) {
	return u.GetUserByField("email", email)
}

func (u *User) GetUserByAccount(account string) (*User, error) {
	return u.GetUserByField("account", account)
}

func (u *User) Get() (*User, error) {
	users := &User{}
	if err := DB.Model(users).First(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
