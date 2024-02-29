package model

type Role struct {
	Role_id   int    `gorm:"primarykey;column:role_id"`
	Role_name string `gorm:"column:role_name"`

	User []User `gorm:"foreign_key:role_id;references:role_id"`
}

func (r *Role) TableName() string {
	return "role"
}
