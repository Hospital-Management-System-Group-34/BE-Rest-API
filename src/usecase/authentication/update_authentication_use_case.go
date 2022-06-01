package authentication

import (
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type updateAuthenticationUseCase struct {
	jwtTokenManager          application.TokenManager
	authenticationRepostiory domain.AuthenticationRepository
}

func NewUpdateAuthenticationUseCase(
	jwtTokenManager application.TokenManager,
	authenticationRepository domain.AuthenticationRepository,
) domain.UpdateAuthenticationUseCase {
	newUpdateAuhtenticationUseCase := updateAuthenticationUseCase{
		jwtTokenManager:          jwtTokenManager,
		authenticationRepostiory: authenticationRepository,
	}

	return &newUpdateAuhtenticationUseCase
}

func (u *updateAuthenticationUseCase) Execute(payload entity.RefreshTokenPayload) (
	entity.NewAccessToken, int, error,
) {
	if err := u.jwtTokenManager.VerifyRefreshToken(payload.RefreshToken); err != nil {
		return entity.NewAccessToken{}, http.StatusBadRequest, err
	}

	if code, err := u.authenticationRepostiory.VerifyRefreshTokenExistence(payload); err != nil {
		return entity.NewAccessToken{}, code, err
	}

	decodedPayload, err := u.jwtTokenManager.DecodeRefreshTokenPayload(payload.RefreshToken)
	if err != nil {
		return entity.NewAccessToken{}, http.StatusInternalServerError, err
	}

	accessToken := u.jwtTokenManager.GenerateAccessToken(decodedPayload)
	newAccessToken := entity.NewAccessToken{
		AccessToken: accessToken,
	}

	return newAccessToken, http.StatusOK, nil
}
