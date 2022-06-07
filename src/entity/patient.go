package entity

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	FName     string `gorm:"not null;column:fname" json:"fName" validate:"required"`
	LName     string `gorm:"not null;column:lname" json:"lName" validate:"required"`
	Phone     string `gorm:"not null" json:"phone" validate:"required"`
	Complaint string `gorm:"not null" json:"complaint" validate:"required"`
}

type UpdatePatientPayload struct {
	ID        uint   `param:"patientID" validate:"required,number,gt=0"`
	FName     string `json:"fName" validate:"required"`
	LName     string `json:"lName" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Complaint string `json:"complaint" validate:"required"`
}

type PatientIDPayload struct {
	ID uint `param:"patientID" validate:"required,number,gt=0"`
}
