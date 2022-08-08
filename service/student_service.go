package service

import (
	"students/dto"
	"students/model"
	"students/repository"
)

type StudentService interface {
	GetAll() []dto.StudentResponse
	GetById(id int) dto.StudentResponse
	Delete(id int)
	Create(student dto.StudentCreateRequest) dto.StudentResponse
	Update(id int, request dto.StudentUpdateRequest) dto.StudentResponse
}

type studentService struct {
	studentRepository repository.StudentRepository
}

func NewStudentService(studentRepository repository.StudentRepository) StudentService {
	return &studentService{studentRepository}
}

func (service *studentService) GetAll() []dto.StudentResponse {
	students := service.studentRepository.GetAll()
	var studentResponses []dto.StudentResponse
	for _, student := range students {
		studentResponses = append(studentResponses, dto.StudentResponse{
			Id:          student.Id,
			Name:        student.Name,
			Major:       student.Major,
			DateOfBirth: student.DateOfBirth,
			Address:     student.Address,
		})
	}
	return studentResponses
}

func (service *studentService) GetById(id int) dto.StudentResponse {
	student, err := service.studentRepository.GetById(id)
	if err != nil {
		return dto.StudentResponse{}
	}
	return dto.StudentResponse{
		Id:          student.Id,
		Name:        student.Name,
		Major:       student.Major,
		DateOfBirth: student.DateOfBirth,
		Address:     student.Address,
	}

}

func (service *studentService) Delete(id int) {
	service.studentRepository.Delete(id)
}

func (service *studentService) Create(request dto.StudentCreateRequest) dto.StudentResponse {
	student := model.Student{}
	// studentResponse := dto.StudentResponse{}

	student.Name = request.Name
	student.Major = request.Major
	student.DateOfBirth = request.DateOfBirth
	student.Address = request.Address
	student = service.studentRepository.Create(student)

	return dto.StudentResponse{
		Id:          student.Id,
		Name:        student.Name,
		Major:       student.Major,
		DateOfBirth: student.DateOfBirth,
		Address:     student.Address,
	}

}

func (service *studentService) Update(id int, request dto.StudentUpdateRequest) dto.StudentResponse {

	student := model.Student{}

	student.Name = request.Name
	student.Major = request.Major
	student.DateOfBirth = request.DateOfBirth
	student.Address = request.Address
	student = service.studentRepository.Update(id, student)

	return dto.StudentResponse{
		Id:          student.Id,
		Name:        student.Name,
		Major:       student.Major,
		DateOfBirth: student.DateOfBirth,
		Address:     student.Address,
	}

}
