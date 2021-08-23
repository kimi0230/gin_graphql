package validateparam

// 驗證格式struct
type TokenValid struct {
	UserToken string `json:"userToken" form:"userToken"  binding:"required,min=20"`
	Lang      string `json:"lang" form:"lang" binding:"-"`
	UserZone  string `json:"user_zone" form:"user_zone" binding:"-"`
}

type SinginValid struct {
	ThirdID  string `json:"third_id" form:"third_id" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required" `
	Name     string `json:"name" form:"name" binding:"required"`
	Sex      string `json:"sex" form:"sex"`
	Nation   string `json:"nation" form:"nation"`
}

type SignoutValid struct {
	TokenValid
}
