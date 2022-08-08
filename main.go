package main

import (
	"students/config"
	"students/controller"
	"students/repository"
	"students/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.NewDB()

	// * wiring
	studentRepository := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepository)
	studentController := controller.NewStudentController(studentService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router := gin.Default()
	// router.GET("/students", config.GetAllStudent)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/students", studentController.GetAllStudent)
	router.GET("/students/:id", studentController.GetStudentById)
	router.DELETE("/students/:id", studentController.DeleteStudent)
	router.POST("/students", studentController.AddStudent)
	router.PUT("/students/:id", studentController.UpdateStudent)

	router.POST("/register", userController.RegisterUser)

	router.Run("localhost:8080")
}
