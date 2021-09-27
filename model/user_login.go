package model

type UserLogin struct {
	ID       int64  `gorm:"id" json:"id"`
	UserName string `gorm:"user_name" json:"user_name"`
	Password string `gorm:"password" json:"password"`
}

func (UserLogin) TableName() string {
	return "user_login"
}
