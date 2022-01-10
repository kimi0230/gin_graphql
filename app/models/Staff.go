package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Staff struct {
	BaseModelIDInt64
	Name     string `json:"name" gorm:"Column:name;type:varchar(20);not null;comment:'名字' "`
	Email    string `json:"email" gorm:"Column:email;type:varchar(50);not null;comment:'信箱' "`
	Password string `json:"password" gorm:"Column:password;type:varchar(50);not null;comment:'密碼' "`
}

func (Staff) TableName() string {
	return "staffs"
}

// 產生密碼
func (s *Staff) HashPassword(password string) error {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	s.Password = string(passwordHash)

	return nil
}

// 比對密碼是否正確
func (s *Staff) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(s.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (s *Staff) Create(user *User) (*User, error) {
	if err := DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Staff) GetUserByField(field, value string) (*User, error) {
	user := &User{}
	if err := DB.Model(user).Where(fmt.Sprintf("%v = ?", field), value).Take(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (s *Staff) GetUserByID(id int) (*User, error) {
	user := &User{}
	if err := DB.Model(user).Where("id = ?", id).First(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (s *Staff) GetUserByEmail(email string) (*User, error) {
	return s.GetUserByField("email", email)
}

func (s *Staff) GetUserByAccount(account string) (*User, error) {
	return s.GetUserByField("account", account)
}

func (s *Staff) Get() (*User, error) {
	users := &User{}
	if err := DB.Model(users).First(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
