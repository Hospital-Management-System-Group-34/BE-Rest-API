package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
	"github.com/golang-jwt/jwt"
)

type jwtTokenManager struct {
}

func NewJWTTokenManager() application.TokenManager {
	newJWTTokenManager := jwtTokenManager{}

	return &newJWTTokenManager
}

func (j *jwtTokenManager) GenerateRefreshToken(payload entity.AuthenticationPayload) string {
	refreshTokenAge, _ := time.ParseDuration(os.Getenv(("REFRESH_TOKEN_AGE")))
	refreshTokenSecretKey := os.Getenv("REFRESH_TOKEN_KEY")

	return j.generateToken(payload.ID, refreshTokenAge, refreshTokenSecretKey)
}

func (j *jwtTokenManager) GenerateAccessToken(payload entity.AuthenticationPayload) string {
	accessTokenAge, _ := time.ParseDuration(os.Getenv(("ACCESS_TOKEN_AGE")))
	accessTokenSecretKey := os.Getenv("ACCESS_TOKEN_KEY")

	return j.generateToken(payload.ID, accessTokenAge, accessTokenSecretKey)
}

func (j *jwtTokenManager) VerifyRefreshToken(refreshToken string) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_KEY")), nil
	})
	if err != nil {
		return fmt.Errorf("invalid refresh token")
	}

	return nil
}

func (j *jwtTokenManager) DecodeRefreshTokenPayload(refreshToken string) (
	entity.AuthenticationPayload, error,
) {
	return j.decodePayload(refreshToken, os.Getenv("REFRESH_TOKEN_KEY"))
}

func (j *jwtTokenManager) DecodeAccessTokenPayload(accessToken string) (
	entity.AuthenticationPayload, error,
) {
	return j.decodePayload(accessToken, os.Getenv("ACCESS_TOKEN_KEY"))
}

func (j *jwtTokenManager) decodePayload(token string, secretKey string) (
	entity.AuthenticationPayload, error,
) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return entity.AuthenticationPayload{}, err
	}

	authenticationPayload := entity.AuthenticationPayload{
		ID: uint(claims["id"].(float64)),
	}
	return authenticationPayload, nil
}

func (j *jwtTokenManager) generateToken(id uint, expirationTime time.Duration, secretKey string) string {
	claims := &Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secretKey))

	return tokenString
}
