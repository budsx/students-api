package model

type User struct {
	Id       int    `gorm:"column:id;primary_key" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}
