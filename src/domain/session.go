package domain

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/labstack/echo/v4"
)

type SessionHandler interface {
	PostSessionHandler(c echo.Context) error
	// GetSessionsHandler(c echo.Context) error
}

type SessionRepository interface {
	AddSession(payload entity.Session) (entity.Session, int, error)
	GetSessionLastQueue(scheduleID uint) (int, int, error)
	// GetSessions() ([]entity.Session, int, error)
}

type AddSessionUseCase interface {
	Execute(payload entity.Session) (entity.Session, int, error)
}

type GetSessionsUseCase interface {
	Execute() (int, error)
}
