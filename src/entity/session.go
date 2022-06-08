package entity

type Session struct {
	PatientID  uint `gorm:"not null" json:"patientID" validate:"required,number,gt=0"`
	ClinicID   uint `gorm:"not null" json:"clinicID" validate:"required,number,gt=0"`
	DoctorID   uint `gorm:"not null" json:"doctorID" validate:"required,number,gt=0"`
	ScheduleID uint `gorm:"not null" json:"scheduleID" validate:"required,number,gt=0"`
	Queue      int  `gorm:"not null" json:"queue"`
}
