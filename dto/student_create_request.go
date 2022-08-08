package dto

type StudentCreateRequest struct {
	Name        string `json:"name"`
	Major       string `json:"major"`
	DateOfBirth string `json:"date_of_birth"`
	Address     string `json:"address"`
}
