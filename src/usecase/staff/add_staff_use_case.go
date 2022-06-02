package staff

import (
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type addStaffUseCase struct {
	staffRepository    domain.StaffRepository
	bcryptPasswordHash application.PasswordHash
}

func NewAddStaffUseCase(
	staffRepository domain.StaffRepository,
	bcryptPasswordHash application.PasswordHash,
) domain.AddStaffUseCase {
	newAddStaffUseCase := addStaffUseCase{
		staffRepository:    staffRepository,
		bcryptPasswordHash: bcryptPasswordHash,
	}

	return &newAddStaffUseCase
}

func (u *addStaffUseCase) Execute(payload entity.Staff) (int, error) {
	if code, err := u.staffRepository.VerifyEmailAvailable(payload.Email); err != nil {
		return code, err
	}

	hashedPassword, err := u.bcryptPasswordHash.Hash(payload.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	payload.Password = hashedPassword

	return u.staffRepository.AddStaff(payload)
}
