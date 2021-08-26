package models

type User struct {
	BaseModel
	Account  string `json:"account"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Get() (*User, error) {
	users := &User{}
	if err := db.Model(users).First(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
