package entity

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	FName    string `gorm:"not null;column:fname" json:"fName" validate:"required"`
	LName    string `gorm:"not null;column:lname" json:"lName" validate:"required"`
	Phone    string `gorm:"not null" json:"phone" validate:"required"`
	ClinicID uint   `gorm:"not null" json:"clinic_id" validate:"required,number,gt=0"`
}

type UpdateDoctorPayload struct {
	ID       uint   `param:"doctorID" validate:"required,number,gt=0"`
	FName    string `json:"fName" validate:"required"`
	LName    string `json:"lName" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	ClinicID uint   `json:"clinic_id" validate:"required,number,gt=0"`
}

type DoctorIDPayload struct {
	ID uint `param:"doctorID" validate:"required,number,gt=0"`
}
