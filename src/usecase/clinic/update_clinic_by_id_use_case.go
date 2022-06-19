package clinic

import (
	// "fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type updateClinicByIDUseCase struct {
	clinicRepository domain.ClinicRepository
	jwtTokenManager  application.TokenManager
	staffRepository  domain.StaffRepository
}

func NewUpdateClinicByIDUseCase(
	clinicRepository domain.ClinicRepository,
	jwtTokenManager application.TokenManager,
	staffRepository domain.StaffRepository,
) domain.UpdateClinicByIDUseCase {
	return &updateClinicByIDUseCase{
		clinicRepository: clinicRepository,
		jwtTokenManager:  jwtTokenManager,
		staffRepository:  staffRepository,
	}
}

func (u *updateClinicByIDUseCase) Execute(
	payload entity.UpdateClinicPayload,
	authorizationHeader entity.AuthorizationHeader,
) (int, error) {
	return http.StatusOK, nil

	// decodedPayload, code, err := u.jwtTokenManager.DecodeAccessTokenPayload(authorizationHeader.AccessToken)
	// if err != nil {
	// 	return code, err
	// }

	// staff, code, err := u.staffRepository.GetStaffByID(decodedPayload.ID)
	// if err != nil {
	// 	return code, err
	// }

	// if staff.StaffType != "admin" {
	// 	return http.StatusForbidden, fmt.Errorf("restricted resource")
	// }

	return u.clinicRepository.UpdateClinicByID(payload)
}
