package service

import (
	"students/dto"
	"students/model"
	"students/repository"
	"students/utils"
)

type UserService interface {
	RegisterUser(request dto.UserRegisterRequest)
}

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{userRepository}
}

func (service *userServiceImpl) RegisterUser(request dto.UserRegisterRequest) {
	user := model.User{}

	hashedPassword := utils.HashAndSalt([]byte(request.Password))
	user.Name = request.Name
	user.Username = request.Username
	user.Password = hashedPassword

	service.userRepository.Register(user)
}
