package repository

import (
	"students/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user model.User) model.User
	Login(username string, password string)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (repo *userRepositoryImpl) Register(user model.User) model.User {
	repo.db.Create(&user)
	return user
}

func (repo *userRepositoryImpl) Login(username string, password string) {
	repo.db.Where("username = ?", username).First(&model.User{})
}
