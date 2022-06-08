package authentication

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type staffLoginUseCase struct {
	staffRepository          domain.StaffRepository
	bcryptPasswordHash       application.PasswordHash
	jwtTokenManager          application.TokenManager
	authenticationRepository domain.AuthenticationRepository
}

func NewStaffLoginUseCase(
	staffRepository domain.StaffRepository,
	bcryptPasswordHash application.PasswordHash,
	jwtTokenManager application.TokenManager,
	authenticationRepository domain.AuthenticationRepository,
) domain.StaffLoginUseCase {
	return &staffLoginUseCase{
		staffRepository:          staffRepository,
		bcryptPasswordHash:       bcryptPasswordHash,
		jwtTokenManager:          jwtTokenManager,
		authenticationRepository: authenticationRepository,
	}
}

func (u *staffLoginUseCase) Execute(payload entity.LoginPayload) (entity.NewLogin, int, error) {
	staff, code, err := u.staffRepository.GetStaffByEmail(payload.Email)
	if err != nil {
		return entity.NewLogin{}, code, err
	}

	if code, err = u.bcryptPasswordHash.ComparePassword(payload.Password, staff.Password); err != nil {
		return entity.NewLogin{}, http.StatusBadRequest, fmt.Errorf("invalid credential")
	}

	authenticationPayload := entity.AuthenticationPayload{
		ID: staff.ID,
	}

	refreshToken, code, err := u.jwtTokenManager.GenerateRefreshToken(authenticationPayload)
	if err != nil {
		return entity.NewLogin{}, code, err
	}

	accessToken, code, err := u.jwtTokenManager.GenerateAccessToken(authenticationPayload)
	if err != nil {
		return entity.NewLogin{}, code, err
	}

	authentication := entity.Authentication{
		Token: refreshToken,
	}
	if code, err := u.authenticationRepository.AddRefreshToken(authentication); err != nil {
		return entity.NewLogin{}, code, err
	}

	newLogin := entity.NewLogin{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return newLogin, http.StatusOK, nil
}
