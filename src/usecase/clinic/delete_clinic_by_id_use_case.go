package clinic

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type deleteClinicByIDUseCase struct {
	clinicRepository domain.ClinicRepository
	jwtTokenManager  application.TokenManager
	staffRepository  domain.StaffRepository
}

func NewDeleteClinicByIDUseCase(
	clinicRepository domain.ClinicRepository,
	jwtTokenManager application.TokenManager,
	staffRepository domain.StaffRepository,
) domain.DeleteClinicByIDUseCase {
	return &deleteClinicByIDUseCase{
		clinicRepository: clinicRepository,
		jwtTokenManager:  jwtTokenManager,
		staffRepository:  staffRepository,
	}
}

func (u *deleteClinicByIDUseCase) Execute(
	id uint,
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

	return u.clinicRepository.DeleteClinicByID(id)
}
