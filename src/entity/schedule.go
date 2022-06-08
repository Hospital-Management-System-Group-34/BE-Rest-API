package entity

import "gorm.io/gorm"

type Schedule struct {
	gorm.Model
	Day      string `gorm:"not null" json:"day" validate:"required"`
	Time     string `gorm:"not null" json:"time" validate:"required"`
	DoctorID uint   `gorm:"not null" json:"doctorID" validate:"required,number,gt=0"`
}

type UpdateSchedulePayload struct {
	ID       uint   `param:"scheduleID" validate:"required,number,gt=0"`
	Day      string `json:"day" validate:"required"`
	Time     string `json:"time" validate:"required"`
	DoctorID uint   `json:"doctorID" validate:"required,number,gt=0"`
}

type ScheduleIDPayload struct {
	ID uint `param:"scheduleID" validate:"required,number,gt=0"`
}
