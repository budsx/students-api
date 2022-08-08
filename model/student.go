package model

type Student struct {
	Id          int    `gorm:"column:student_id;primary_key" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Major       string `gorm:"column:major" json:"major"`
	DateOfBirth string `gorm:"column:date_of_birth" json:"date_of_birth"`
	Address     string `gorm:"column:address" json:"address"`
}
