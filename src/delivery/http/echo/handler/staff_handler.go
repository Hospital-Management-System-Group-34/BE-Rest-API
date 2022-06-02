package handler

import (
	"log"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/util"

	"github.com/labstack/echo/v4"
)

type staffHandler struct {
	addStaffUseCase domain.AddStaffUseCase
}

func NewStaffHandler(addStaffUseCase domain.AddStaffUseCase) domain.StaffHandler {
	newStaffHandler := staffHandler{
		addStaffUseCase: addStaffUseCase,
	}

	return &newStaffHandler
}

func (h *staffHandler) PostStaffHandler(c echo.Context) error {
	payload := entity.Staff{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	code, err := h.addStaffUseCase.Execute(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}
