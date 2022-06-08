package routes

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/handler"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/middleware"
	repository "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/repository/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/jwt"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/clinic"
	"github.com/labstack/echo/v4"
)

func ClinicRoutes(e *echo.Echo) {
	postgresDB := postgres.Connect()

	clinicRepository := repository.NewClinicRepository(postgresDB)
	jwtTokenManager := jwt.NewJWTTokenManager()
	staffRepository := repository.NewStaffRepository(postgresDB)

	addClinicUseCase := clinic.NewAddClinicUseCase(clinicRepository, jwtTokenManager, staffRepository)
	getClinicsUseCase := clinic.NewGetClinicsUseCase(clinicRepository)
	getClinicByIDUseCase := clinic.NewGetClinicByIDUseCase(clinicRepository)
	updateClinicByIDUseCase := clinic.NewUpdateClinicByIDUseCase(
		clinicRepository,
		jwtTokenManager,
		staffRepository,
	)
	deleteClinicByIDUseCase := clinic.NewDeleteClinicByIDUseCase(
		clinicRepository,
		jwtTokenManager,
		staffRepository,
	)

	clinicHandler := handler.NewClinicHandler(
		addClinicUseCase,
		getClinicsUseCase,
		getClinicByIDUseCase,
		updateClinicByIDUseCase,
		deleteClinicByIDUseCase,
	)

	e.POST("/clinics", clinicHandler.PostClinicHandler, middleware.JWTMiddleware())
	e.GET("/clinics", clinicHandler.GetClinicsHandler, middleware.JWTMiddleware())
	e.GET("/clinics/:clinicID", clinicHandler.GetClinicByIDHandler, middleware.JWTMiddleware())
	e.PUT("/clinics/:clinicID", clinicHandler.PutClinicByIDHandler, middleware.JWTMiddleware())
	e.DELETE("/clinics/:clinicID", clinicHandler.DeleteClinicByIDHandler, middleware.JWTMiddleware())
}
