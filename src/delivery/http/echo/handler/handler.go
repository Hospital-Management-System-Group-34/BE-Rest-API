package handler

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/util"
	"github.com/labstack/echo/v4"
)

type helloHandler struct {
	helloUseCase domain.HelloUseCase
}

func NewHelloHandler(h domain.HelloUseCase) domain.HelloHandler {
	newHelloHandler := helloHandler{
		helloUseCase: h,
	}

	return &newHelloHandler
}

func (h *helloHandler) GetHelloHandler(c echo.Context) error {
	data := h.helloUseCase.Execute()

	return c.JSON(util.SuccessResponseWithData(data))
}
