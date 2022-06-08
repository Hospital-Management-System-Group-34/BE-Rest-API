package server

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/routes"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/validator"

	"github.com/labstack/echo/v4"
)

func CreateServer() *echo.Echo {
	e := echo.New()

	routes.HelloRoutes(e)
	routes.StaffRoutes(e)
	routes.AuthenticationRoutes(e)
	routes.PatientRoutes(e)
	routes.ClinicRoutes(e)
	routes.DoctorRoutes(e)
	routes.ScheduleRoutes(e)

	validator.NewValidator(e)

	return e
}
