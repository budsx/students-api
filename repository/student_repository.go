package repository

import (
	"students/exception"
	"students/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	GetAll() []model.Student
	GetById(id int) (model.Student, *exception.AppError)
	Delete(id int)
	Create(student model.Student) model.Student
	Update(id int, student model.Student) model.Student
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db}
}

func (repo *studentRepository) GetAll() []model.Student {
	var students []model.Student
	repo.db.Find(&students)
	return students
}

func (repo *studentRepository) GetById(id int) (model.Student, *exception.AppError) {
	var student model.Student
	repo.db.First(&student, id)
	if student.Id == 0 {
		return student, exception.NotFoundError("Student not found")
	} else {
		return student, nil
	}
}

func (repo *studentRepository) Delete(id int) {
	repo.db.Delete(&model.Student{}, id)
}

func (repo *studentRepository) Create(student model.Student) model.Student {
	repo.db.Create(&student)
	return student
}

func (repo *studentRepository) Update(id int, student model.Student) model.Student {
	repo.db.Model(&student).Where("student_id = ?", id).Updates(&student)
	return student
}
