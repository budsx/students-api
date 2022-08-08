package controller

import (
	"strconv"
	"students/dto"
	"students/service"

	"github.com/gin-gonic/gin"
)

type StudentController interface {
	GetAllStudent(c *gin.Context)
	GetStudentById(c *gin.Context)
	DeleteStudent(c *gin.Context)
	AddStudent(c *gin.Context)
	UpdateStudent(c *gin.Context)
}

type studentController struct {
	studentService service.StudentService
}

func NewStudentController(studentService service.StudentService) StudentController {
	return &studentController{studentService}
}

func (controller *studentController) GetAllStudent(c *gin.Context) {
	students := controller.studentService.GetAll()
	c.JSON(200, gin.H{
		"students": students,
	})
}

func (controller *studentController) GetStudentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	student := controller.studentService.GetById(id)
	c.JSON(200, gin.H{
		"data": student,
	})
}

func (controller *studentController) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	controller.studentService.Delete(id)
}

func (controller *studentController) AddStudent(c *gin.Context) {
	student := dto.StudentCreateRequest{}
	c.BindJSON(&student)
	controller.studentService.Create(student)
}

func (controller *studentController) UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	student := dto.StudentUpdateRequest{}
	c.BindJSON(&student)
	controller.studentService.Update(id, student)

}
