package handler

import (
	"log"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/util"
	"github.com/labstack/echo/v4"
)

type doctorHandler struct {
	addDoctorUseCase        domain.AddDoctorUseCase
	getDoctorsUseCase       domain.GetDoctorsUseCase
	getDoctorByIDUseCase    domain.GetDoctorByIDUseCase
	updateDoctorByIDUseCase domain.UpdateDoctorByIDUseCase
	deleteDoctorByIDUseCase domain.DeleteDoctorByIDUseCase
}

func NewDoctorHandler(
	addDoctorUseCase domain.AddDoctorUseCase,
	getDoctorsUseCase domain.GetDoctorsUseCase,
	getDoctorByIDUseCase domain.GetDoctorByIDUseCase,
	updateDoctorByIDUseCase domain.UpdateDoctorByIDUseCase,
	deleteDoctorByIDUseCase domain.DeleteDoctorByIDUseCase,
) domain.DoctorHandler {
	return &doctorHandler{
		addDoctorUseCase:        addDoctorUseCase,
		getDoctorsUseCase:       getDoctorsUseCase,
		getDoctorByIDUseCase:    getDoctorByIDUseCase,
		updateDoctorByIDUseCase: updateDoctorByIDUseCase,
		deleteDoctorByIDUseCase: deleteDoctorByIDUseCase,
	}
}

func (h *doctorHandler) PostDoctorHandler(c echo.Context) error {
	payload := entity.Doctor{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	code, err := h.addDoctorUseCase.Execute(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *doctorHandler) GetDoctorsHandler(c echo.Context) error {
	doctors, code, err := h.getDoctorsUseCase.Execute()
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(doctors))
}

func (h *doctorHandler) GetDoctorByIDHandler(c echo.Context) error {
	payload := entity.DoctorIDPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	doctor, code, err := h.getDoctorByIDUseCase.Execute(payload.ID)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(doctor))
}

func (h *doctorHandler) PutDoctorByIDHandler(c echo.Context) error {
	payload := entity.UpdateDoctorPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	code, err := h.updateDoctorByIDUseCase.Execute(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *doctorHandler) DeleteDoctorByIDHandler(c echo.Context) error {
	payload := entity.DoctorIDPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	code, err := h.deleteDoctorByIDUseCase.Execute(payload.ID)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}