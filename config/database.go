package config

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// * postgres
	db, err := gorm.Open(postgres.Open("postgres://postgres:218799@localhost:5432/go-project?sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
