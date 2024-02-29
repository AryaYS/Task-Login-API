package model

type User struct {
	User_name string `gorm:"primary_key;column:user_name"`
	User_pass string `gorm:"column:user_pass"`
	Role_id   int    `gorm:"column:role_id"`
}

func (u *User) TableName() string {
	return "user"
}
