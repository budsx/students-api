package controller

import (
	"students/dto"
	"students/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(c *gin.Context)
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
	controller.userService.RegisterUser(user)
}
