package model

import "github.com/golang-jwt/jwt/v4"

type User struct {
	Id       int    `gorm:"column:id;primary_key" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}

type UserLogin struct {
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type AccessTokenClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
