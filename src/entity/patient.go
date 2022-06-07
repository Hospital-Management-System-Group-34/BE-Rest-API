package entity

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	FName   string `gorm:"not null;column:fname" json:"fName" validate:"required"`
	LName   string `gorm:"not null;column:lname" json:"lName" validate:"required"`
	Phone   string `gorm:"not null" json:"phone" validate:"required"`
	Keluhan string `gorm:"not null" json:"keluhan" validate:"required"`
}

type UpdatePatientPayload struct {
	ID      uint   `param:"patientID" validate:"required"`
	FName   string `json:"fName" validate:"required"`
	LName   string `json:"lName" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Keluhan string `json:"keluhan" validate:"required"`
}

type PatientIDPayload struct {
	ID uint `param:"patientID" validate:"required"`
}
