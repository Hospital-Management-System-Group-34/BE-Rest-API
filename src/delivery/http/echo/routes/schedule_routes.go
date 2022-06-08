package routes

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/handler"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/middleware"
	repository "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/repository/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/schedule"

	"github.com/labstack/echo/v4"
)

func ScheduleRoutes(e *echo.Echo) {
	postgresDB := postgres.Connect()
	scheduleRepository := repository.NewScheduleRepository(postgresDB)
	doctorRepository := repository.NewDoctorRepository(postgresDB)

	addScheduleUseCase := schedule.NewAddScheduleUseCase(scheduleRepository, doctorRepository)
	getSchedulesUseCase := schedule.NewGetSchedulesUseCase(scheduleRepository)
	getScheduleByIDUseCase := schedule.NewGetScheduleByIDUseCase(scheduleRepository)
	updateScheduleByIDUseCase := schedule.NewUpdateScheduleByIDUseCase(scheduleRepository, doctorRepository)
	deleteScheduleByIDUseCase := schedule.NewDeleteScheduleByIDUseCase(scheduleRepository)

	scheduleHandler := handler.NewScheduleHandler(
		addScheduleUseCase,
		getSchedulesUseCase,
		getScheduleByIDUseCase,
		updateScheduleByIDUseCase,
		deleteScheduleByIDUseCase,
	)

	e.POST("/schedules", scheduleHandler.PostScheduleHandler, middleware.JWTMiddleware())
	e.GET("/schedules", scheduleHandler.GetSchedulesHandler, middleware.JWTMiddleware())
	e.GET("/schedules/:scheduleID", scheduleHandler.GetScheduleByIDHandler, middleware.JWTMiddleware())
	e.PUT("/schedules/:scheduleID", scheduleHandler.PutScheduleByIDHandler, middleware.JWTMiddleware())
	e.DELETE("/schedules/:scheduleID", scheduleHandler.DeleteScheduleByIDHandler, middleware.JWTMiddleware())
}
