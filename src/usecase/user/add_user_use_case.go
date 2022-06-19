package user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
	"github.com/aidarkhanov/nanoid"
)

type addUserUseCase struct {
	userRepository     domain.UserRepository
	bcryptPasswordHash application.PasswordHash
	jwtTokenManager    application.TokenManager
}

func NewAddUserUseCase(
	userRepository domain.UserRepository,
	bcryptPasswordHash application.PasswordHash,
	jwtTokenManager application.TokenManager,
) domain.AddUserUseCase {
	return &addUserUseCase{
		userRepository:     userRepository,
		bcryptPasswordHash: bcryptPasswordHash,
		jwtTokenManager:    jwtTokenManager,
	}
}

func (u *addUserUseCase) Execute(
	payload entity.User,
	authorizationHeader entity.AuthorizationHeader,
) (entity.AddedUser, int, error) {
	decodedPayload, code, err := u.jwtTokenManager.DecodeAccessTokenPayload(authorizationHeader.AccessToken)
	if err != nil {
		return entity.AddedUser{}, code, err
	}

	user, code, err := u.userRepository.GetUserByID(decodedPayload.ID)
	if err != nil {
		return entity.AddedUser{}, code, err
	}

	if user.Role != "Admin" {
		return entity.AddedUser{}, http.StatusForbidden, fmt.Errorf("restricted resource")
	}

	if payload.Role == "Admin" {
		return entity.AddedUser{}, http.StatusBadRequest, fmt.Errorf("admin user already exists")
	}

	if payload.Role != "Doctor" && payload.Role == "Staff" {
		return entity.AddedUser{}, http.StatusBadRequest, fmt.Errorf("role must be Doctor or Staff")
	}

	nanoid, err := nanoid.Generate(nanoid.DefaultAlphabet, 5)
	if err != nil {
		return entity.AddedUser{}, http.StatusInternalServerError, err
	}
	payload.ID = fmt.Sprintf("%s-%s", strings.ToLower(payload.Role), strings.ToLower(nanoid))

	hashedPassword, code, err := u.bcryptPasswordHash.Hash(payload.Password)
	if err != nil {
		return entity.AddedUser{}, code, err
	}

	payload.Password = hashedPassword

	return u.userRepository.AddUser(payload)
}
