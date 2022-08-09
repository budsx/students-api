package main

import (
	"students/config"
	"students/controller"
	"students/middleware"
	"students/repository"
	"students/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db := config.NewDB()
	godotenv.Load()

	// * wiring
	studentRepository := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepository)
	studentController := controller.NewStudentController(studentService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	// authService := service.NewAuthService(userRepository)
	// authController := controller.NewAuthController(authService)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	students := router.Group("/students").Use(middleware.AuthMiddleware)
	// students.Use(middleware.AuthMiddleware)

	{
		students.GET("", studentController.GetAllStudent)
		students.GET("/:id", studentController.GetStudentById)
		students.DELETE("/:id", studentController.DeleteStudent)
		students.POST("", studentController.AddStudent)
		students.PUT("/:id", studentController.UpdateStudent)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/login", userController.LoginUser)
		auth.POST("/register", userController.RegisterUser)
	}

	router.Run("localhost:8080")
}
