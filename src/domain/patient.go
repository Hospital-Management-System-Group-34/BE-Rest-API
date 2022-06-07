package domain

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/labstack/echo/v4"
)

type PatientHandler interface {
	PostPatientHandler(c echo.Context) error
	GetPatientsHandler(c echo.Context) error
	GetPatientByIDHandler(c echo.Context) error
	PutPatientByIDHandler(c echo.Context) error
	DeletePatientByIDHandler(c echo.Context) error
}

type PatientRepository interface {
	AddPatient(payload entity.Patient) (int, error)
	GetPatients() ([]entity.Patient, int, error)
	GetPatientByID(id uint) (entity.Patient, int, error)
	UpdatePatientByID(payload entity.UpdatePatientPayload) (int, error)
	DeletePatientByID(id uint) (int, error)
}

type AddPatientUseCase interface {
	Execute(payload entity.Patient) (int, error)
}

type GetPatientsUseCase interface {
	Execute() ([]entity.Patient, int, error)
}

type GetPatientByIDUseCase interface {
	Execute(id uint) (entity.Patient, int, error)
}

type UpdatePatientByIDUseCase interface {
	Execute(payload entity.UpdatePatientPayload) (int, error)
}

type DeletePatientByIDUseCase interface {
	Execute(id uint) (int, error)
}
