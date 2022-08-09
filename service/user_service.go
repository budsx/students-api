package service

import (
	"errors"
	"log"
	"os"
	"students/dto"
	"students/model"
	"students/repository"
	"students/utils"

	"github.com/golang-jwt/jwt/v4"
)

type UserService interface {
	RegisterUser(request dto.UserRegisterRequest) error
	Login(request dto.UserLoginRequest) (dto.AccessTokenResponse, error)
}

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{userRepository}
}

func (service *userServiceImpl) RegisterUser(request dto.UserRegisterRequest) error {
	user := model.User{}

	hashedPassword := utils.HashPassword([]byte(request.Password))
	user.Name = request.Name
	user.Username = request.Username
	user.Password = hashedPassword

	// * check if user already exists

	users, _ := service.userRepository.FindByUsername(user.Username)
	if users.Username == user.Username {
		log.Println("user already exists")
		return errors.New("user already exists")
	} else {
		service.userRepository.Register(user)
		return nil
	}

}

func (service *userServiceImpl) Login(request dto.UserLoginRequest) (dto.AccessTokenResponse, error) {
	username := request.Username
	password := request.Password
	users, err := service.userRepository.FindByUsername(username)

	if username == "" || password == "" {
		log.Println("username or password is empty")
		log.Println(err)
	}

	// * compare password
	errs := utils.VerifyPassword(password, users.Password)
	if errs != nil {
		log.Println("password is not correct")
		return dto.AccessTokenResponse{}, errs
	}

	token, _ := GenerateToken(username)

	return dto.AccessTokenResponse{
		AccessToken: token,
	}, err
}

func GenerateToken(username string) (string, error) {
	claims := model.AccessTokenClaims{
		Username: username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Println("error when generate token")
	}
	return tokenString, err
}
