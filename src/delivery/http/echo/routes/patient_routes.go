package routes

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/handler"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/middleware"
	repository "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/repository/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/patient"

	"github.com/labstack/echo/v4"
)

func PatientRoutes(e *echo.Echo) {
	postgresDB := postgres.Connect()
	PatientRepository := repository.NewPatientRepository(postgresDB)

	addPatientUseCase := patient.NewAddPatientUseCase(PatientRepository)
	getPatientsUseCase := patient.NewGetPatientsUseCase(PatientRepository)
	getPatientByIDUseCase := patient.NewGetPatientByIDUseCase(PatientRepository)
	updatePatientByIDUseCase := patient.NewUpdatePatientByIDUseCase(PatientRepository)
	deletePatientByIDUseCase := patient.NewDeletePatientByIDUseCase(PatientRepository)

	PatientHandler := handler.NewPatientHandler(
		addPatientUseCase,
		getPatientsUseCase,
		getPatientByIDUseCase,
		updatePatientByIDUseCase,
		deletePatientByIDUseCase,
	)

	e.POST("/patients", PatientHandler.PostPatientHandler, middleware.JWTMiddleware())
	e.GET("/patients", PatientHandler.GetPatientsHandler, middleware.JWTMiddleware())
	e.GET("/patients/:patientID", PatientHandler.GetPatientByIDHandler, middleware.JWTMiddleware())
	e.PUT("/patients/:patientID", PatientHandler.PutPatientByIDHandler, middleware.JWTMiddleware())
	e.DELETE("/patients/:patientID", PatientHandler.DeletePatientByIDHandler, middleware.JWTMiddleware())
}
