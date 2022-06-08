package handler

import (
	"log"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/util"
	"github.com/labstack/echo/v4"
)

type scheduleHandler struct {
	addScheduleUseCase        domain.AddScheduleUseCase
	getSchedulesUseCase       domain.GetSchedulesUseCase
	getScheduleByIDUseCase    domain.GetScheduleByIDUseCase
	updateScheduleByIDUseCase domain.UpdateScheduleByIDUseCase
	deleteScheduleByIDUseCase domain.DeleteScheduleByIDUseCase
}

func NewScheduleHandler(
	addScheduleUseCase domain.AddScheduleUseCase,
	getSchedulesUseCase domain.GetSchedulesUseCase,
	getScheduleByIDUseCase domain.GetScheduleByIDUseCase,
	updateScheduleByIDUseCase domain.UpdateScheduleByIDUseCase,
	deleteScheduleByIDUseCase domain.DeleteScheduleByIDUseCase,
) domain.ScheduleHandler {
	return &scheduleHandler{
		addScheduleUseCase:        addScheduleUseCase,
		getSchedulesUseCase:       getSchedulesUseCase,
		getScheduleByIDUseCase:    getScheduleByIDUseCase,
		updateScheduleByIDUseCase: updateScheduleByIDUseCase,
		deleteScheduleByIDUseCase: deleteScheduleByIDUseCase,
	}
}

func (h *scheduleHandler) PostScheduleHandler(c echo.Context) error {
	payload := entity.Schedule{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	code, err := h.addScheduleUseCase.Execute(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *scheduleHandler) GetSchedulesHandler(c echo.Context) error {
	Schedules, code, err := h.getSchedulesUseCase.Execute()
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(Schedules))
}

func (h *scheduleHandler) GetScheduleByIDHandler(c echo.Context) error {
	payload := entity.ScheduleIDPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	Schedule, code, err := h.getScheduleByIDUseCase.Execute(payload.ID)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(Schedule))
}

func (h *scheduleHandler) PutScheduleByIDHandler(c echo.Context) error {
	payload := entity.UpdateSchedulePayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	code, err := h.updateScheduleByIDUseCase.Execute(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *scheduleHandler) DeleteScheduleByIDHandler(c echo.Context) error {
	payload := entity.ScheduleIDPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	code, err := h.deleteScheduleByIDUseCase.Execute(payload.ID)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}
