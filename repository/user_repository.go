package repository

import (
	"fmt"
	"students/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user model.User) model.User
	FindByUsername(username string) (model.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (repo *userRepositoryImpl) Register(user model.User) model.User {
	repo.db.Create(&user)
	fmt.Println("di repos regiter", user.Username)
	return user
}

func (repo *userRepositoryImpl) FindByUsername(username string) (model.User, error) {
	users := model.User{}
	repo.db.Where("username = ?", username).Take(&users)
	// fmt.Println("di repos", users.Password, users.Username)
	return users, nil

}