package doctor

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type updateDoctorByIDUseCase struct {
	doctorRepository domain.DoctorRepository
	clinicRepository domain.ClinicRepository
	jwtTokenManager  application.TokenManager
	staffRepository  domain.StaffRepository
}

func NewUpdateDoctorByIDUseCase(
	doctorRepository domain.DoctorRepository,
	clinicRepository domain.ClinicRepository,
	jwtTokenManager application.TokenManager,
	staffRepository domain.StaffRepository,
) domain.UpdateDoctorByIDUseCase {
	return &updateDoctorByIDUseCase{
		doctorRepository: doctorRepository,
		clinicRepository: clinicRepository,
		jwtTokenManager:  jwtTokenManager,
		staffRepository:  staffRepository,
	}
}

func (u *updateDoctorByIDUseCase) Execute(
	payload entity.UpdateDoctorPayload,
	authorizationHeader entity.AuthorizationHeader,
) (int, error) {
	decodedPayload, code, err := u.jwtTokenManager.DecodeAccessTokenPayload(authorizationHeader.AccessToken)
	if err != nil {
		return code, err
	}

	staff, code, err := u.staffRepository.GetStaffByID(decodedPayload.ID)
	if err != nil {
		return code, err
	}

	if staff.StaffType != "admin" {
		return http.StatusForbidden, fmt.Errorf("restricted resource")
	}

	if _, code, err := u.clinicRepository.GetClinicByID(payload.ClinicID); err != nil {
		return code, err
	}

	return u.doctorRepository.UpdateDoctorByID(payload)
}
