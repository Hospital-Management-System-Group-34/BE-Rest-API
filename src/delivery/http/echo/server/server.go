package server

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/routes"

	"github.com/labstack/echo/v4"
)

func CreateServer() *echo.Echo {
	e := echo.New()

	routes.HelloRoutes(e)

	return e
}
