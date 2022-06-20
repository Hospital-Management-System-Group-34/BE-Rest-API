package clinic

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type addClinicUseCase struct {
	clinicRepository domain.ClinicRepository
	jwtTokenManager  application.TokenManager
	userRepository   domain.UserRepository
}

func NewAddClinicUseCase(
	clinicRepository domain.ClinicRepository,
	jwtTokenManager application.TokenManager,
	userRepository domain.UserRepository,
) domain.AddClinicUseCase {
	return &addClinicUseCase{
		clinicRepository: clinicRepository,
		jwtTokenManager:  jwtTokenManager,
		userRepository:   userRepository,
	}
}

func (u *addClinicUseCase) Execute(
	payload entity.Clinic,
	authorizationHeader entity.AuthorizationHeader,
) (int, error) {
	decodedPayload, code, err := u.jwtTokenManager.DecodeAccessTokenPayload(authorizationHeader.AccessToken)
	if err != nil {
		return code, err
	}

	user, code, err := u.userRepository.GetUserByID(decodedPayload.ID)
	if err != nil {
		return code, err
	}

	if user.Role != "Admin" {
		return http.StatusForbidden, fmt.Errorf("restricted resource")
	}

	return u.clinicRepository.AddClinic(payload)
}
