package application

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type TokenManager interface {
	GenerateRefreshToken(payload entity.AuthenticationPayload) string
	GenerateAccessToken(payload entity.AuthenticationPayload) string
	VerifyRefreshToken(refreshToken string) error
	DecodeRefreshTokenPayload(refreshToken string) (entity.AuthenticationPayload, error)
	DecodeAccessTokenPayload(accessToken string) (entity.AuthenticationPayload, error)
}
