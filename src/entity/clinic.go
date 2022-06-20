package entity

import (
	"time"
)

type Clinic struct {
	ID        string    `gorm:"not null,primaryKey,index:idx_id" json:"id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	Name      string    `gorm:"not null" json:"name" validate:"required"`
}

type UpdateClinicPayload struct {
	ID   string `param:"clinicID" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type ClinicIDPayload struct {
	ID string `param:"clinicID" validate:"required"`
}
