package routes

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/handler"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/middleware"
	repository "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/repository/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/session"
	"github.com/labstack/echo/v4"
)

func SessionRoutes(e *echo.Echo) {
	postgresDB := postgres.Connect()
	sessionRepository := repository.NewSessionRepository(postgresDB)
	patientRepository := repository.NewPatientRepository(postgresDB)
	doctorRepository := repository.NewDoctorRepository(postgresDB)
	clinicRepository := repository.NewClinicRepository(postgresDB)
	scheduleRepository := repository.NewScheduleRepository(postgresDB)

	addSessionUseCase := session.NewAddSessionUseCase(
		sessionRepository,
		patientRepository,
		clinicRepository,
		doctorRepository,
		scheduleRepository,
	)

	sessionHandler := handler.NewSessionHandler(
		addSessionUseCase,
	)

	e.POST("/sessions", sessionHandler.PostSessionHandler, middleware.JWTMiddleware())
}
