package routes

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/handler"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/middleware"
	repository "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/repository/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/jwt"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/nanoid"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/session"
	"github.com/labstack/echo/v4"
)

func SessionRoutes(e *echo.Echo) {
	postgresDB := postgres.Connect()

	sessionRepository := repository.NewSessionRepository(postgresDB)
	patientRepository := repository.NewPatientRepository(postgresDB)
	userRepository := repository.NewUserRepository(postgresDB)
	clinicRepository := repository.NewClinicRepository(postgresDB)
	scheduleRepository := repository.NewScheduleRepository(postgresDB)

	nanoidIDGenerator := nanoid.NewNanoidIDGenerator()
	jwtTokenManager := jwt.NewJWTTokenManager()

	addSessionUseCase := session.NewAddSessionUseCase(
		sessionRepository,
		patientRepository,
		clinicRepository,
		userRepository,
		scheduleRepository,
		nanoidIDGenerator,
	)
	getSessionsUseCase := session.NewGetSessionsUseCase(
		sessionRepository,
		userRepository,
		jwtTokenManager,
	)

	sessionHandler := handler.NewSessionHandler(
		addSessionUseCase,
		getSessionsUseCase,
	)

	e.POST("/sessions", sessionHandler.PostSessionHandler, middleware.JWTMiddleware())
	e.GET("/sessions", sessionHandler.GetSessionsHandler, middleware.JWTMiddleware())
}
