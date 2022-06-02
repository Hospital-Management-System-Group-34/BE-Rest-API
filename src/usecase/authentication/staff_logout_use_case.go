package authentication

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type staffLogoutUseCase struct {
	authenticationRepository domain.AuthenticationRepository
}

func NewStaffLogoutUseCase(authenticationRepository domain.AuthenticationRepository) domain.StaffLogoutUseCase {
	newStaffLogoutUseCase := staffLogoutUseCase{
		authenticationRepository: authenticationRepository,
	}

	return &newStaffLogoutUseCase
}

func (u *staffLogoutUseCase) Execute(payload entity.RefreshTokenPayload) (int, error) {
	if code, err := u.authenticationRepository.VerifyRefreshTokenExistence(payload); err != nil {
		return code, err
	}

	return u.authenticationRepository.DeleteRefreshToken(payload)
}
