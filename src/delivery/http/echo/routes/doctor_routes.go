package routes

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/handler"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/middleware"
	repository "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/repository/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/doctor"

	"github.com/labstack/echo/v4"
)

func DoctorRoutes(e *echo.Echo) {
	postgresDB := postgres.Connect()
	doctorRepository := repository.NewDoctorRepository(postgresDB)

	addDoctorUseCase := doctor.NewAddDoctorUseCase(doctorRepository)
	getDoctorsUseCase := doctor.NewGetDoctorsUseCase(doctorRepository)
	getDoctorByIDUseCase := doctor.NewGetDoctorByIDUseCase(doctorRepository)
	updateDoctorByIDUseCase := doctor.NewUpdateDoctorByIDUseCase(doctorRepository)
	deleteDoctorByIDUseCase := doctor.NewDeleteDoctorByIDUseCase(doctorRepository)

	doctorHandler := handler.NewDoctorHandler(
		addDoctorUseCase,
		getDoctorsUseCase,
		getDoctorByIDUseCase,
		updateDoctorByIDUseCase,
		deleteDoctorByIDUseCase,
	)

	e.POST("/doctors", doctorHandler.PostDoctorHandler, middleware.JWTMiddleware())
	e.GET("/doctors", doctorHandler.GetDoctorsHandler, middleware.JWTMiddleware())
	e.GET("/doctors/:doctorID", doctorHandler.GetDoctorByIDHandler, middleware.JWTMiddleware())
	e.PUT("/doctors/:doctorID", doctorHandler.PutDoctorByIDHandler, middleware.JWTMiddleware())
	e.DELETE("/doctors/:doctorID", doctorHandler.DeleteDoctorByIDHandler, middleware.JWTMiddleware())
}
