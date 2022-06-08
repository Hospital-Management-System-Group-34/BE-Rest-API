package domain

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/labstack/echo/v4"
)

type ScheduleHandler interface {
	PostScheduleHandler(c echo.Context) error
	GetSchedulesHandler(c echo.Context) error
	GetScheduleByIDHandler(c echo.Context) error
	PutScheduleByIDHandler(c echo.Context) error
	DeleteScheduleByIDHandler(c echo.Context) error
}

type ScheduleRepository interface {
	AddSchedule(payload entity.Schedule) (int, error)
	GetSchedules() ([]entity.Schedule, int, error)
	GetScheduleByID(id uint) (entity.Schedule, int, error)
	UpdateScheduleByID(payload entity.UpdateSchedulePayload) (int, error)
	DeleteScheduleByID(id uint) (int, error)
}

type AddScheduleUseCase interface {
	Execute(payload entity.Schedule) (int, error)
}

type GetSchedulesUseCase interface {
	Execute() ([]entity.Schedule, int, error)
}

type GetScheduleByIDUseCase interface {
	Execute(id uint) (entity.Schedule, int, error)
}

type UpdateScheduleByIDUseCase interface {
	Execute(payload entity.UpdateSchedulePayload) (int, error)
}

type DeleteScheduleByIDUseCase interface {
	Execute(id uint) (int, error)
}
