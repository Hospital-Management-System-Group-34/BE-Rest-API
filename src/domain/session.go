package domain

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/labstack/echo/v4"
)

type SessionHandler interface {
	PostSessionHandler(c echo.Context) error
	GetSessionsHandler(c echo.Context) error
}

type SessionRepository interface {
	AddSession(payload entity.Session) (entity.Session, int, error)
	GetSessionLastQueue(scheduleID string, date string) (int, int, error)
	GetSessionByID(id string) (entity.Session, int, error)
	GetSessionByPatientID(patientID string) ([]entity.Session, int, error)
	GetSessions() ([]entity.Session, int, error)
	GetSessionsByDoctorID(doctorID string) ([]entity.Session, int, error)
	GetQueuedSessionsByDoctorID(doctorID string) ([]entity.Session, int, error)
	GetCompletedSessionsByDoctorID(doctorID string) ([]entity.Session, int, error)
	GetCancelledSessionsByDoctorID(doctorID string) ([]entity.Session, int, error)
}

type AddSessionUseCase interface {
	Execute(payload entity.Session) (entity.Session, int, error)
}

type GetSessionsUseCase interface {
	Execute(
		payload entity.GetSessionParams,
		authorizationHeader entity.AuthorizationHeader,
	) ([]entity.Session, int, error)
}
