package controller

import (
	"students/dto"
	"students/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService}
}

func (controller *userController) RegisterUser(c *gin.Context) {
	user := dto.UserRegisterRequest{}
	c.BindJSON(&user)
	err := controller.userService.RegisterUser(user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "User already exists",
		})
	} else {
		c.JSON(201, gin.H{
			"message": "User registered successfully",
		})
	}
}

func (controller *userController) LoginUser(c *gin.Context) {
	user := dto.UserLoginRequest{}
	c.BindJSON(&user)
	token, err := controller.userService.Login(user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Username and password do not match",
		})
	} else {
		c.JSON(200, gin.H{
			"access_token": token,
		})
	}
}
