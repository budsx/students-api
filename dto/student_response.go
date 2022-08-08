package dto

type StudentResponse struct {
	Id          int    `json:"student_id"`
	Name        string `json:"name"`
	Major       string `json:"major"`
	DateOfBirth string `json:"date_of_birth"`
	Address     string `json:"address"`
}
