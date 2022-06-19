package clinic

import (
	// "fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type addClinicUseCase struct {
	clinicRepository domain.ClinicRepository
	jwtTokenManager  application.TokenManager
	staffRepository  domain.StaffRepository
}

func NewAddClinicUseCase(
	clinicRepository domain.ClinicRepository,
	jwtTokenManager application.TokenManager,
	staffRepository domain.StaffRepository,
) domain.AddClinicUseCase {
	return &addClinicUseCase{
		clinicRepository: clinicRepository,
		jwtTokenManager:  jwtTokenManager,
		staffRepository:  staffRepository,
	}
}

func (u *addClinicUseCase) Execute(
	payload entity.Clinic,
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

	return u.clinicRepository.AddClinic(payload)
}
