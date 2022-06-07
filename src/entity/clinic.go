package entity

import "gorm.io/gorm"

type Clinic struct {
	gorm.Model
	Name string `gorm:"not null" json:"name" validate:"required"`
}

type UpdateClinicPayload struct {
	ID   uint   `param:"ClinicID" validate:"required,number,gt=0"`
	Name string `json:"name" validate:"required"`
}

type ClinicIDPayload struct {
	ID uint `param:"clinicID" validate:"required,number,gt=0"`
}
