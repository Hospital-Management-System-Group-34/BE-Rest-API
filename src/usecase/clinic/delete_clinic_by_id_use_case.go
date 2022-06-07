package clinic

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
)

type deleteClinicByIDUseCase struct {
	clinicRepository domain.ClinicRepository
}

func NewDeleteClinicByIDUseCase(clinicRepository domain.ClinicRepository) domain.DeleteClinicByIDUseCase {
	return &deleteClinicByIDUseCase{
		clinicRepository: clinicRepository,
	}
}

func (u *deleteClinicByIDUseCase) Execute(id uint) (int, error) {
	return u.clinicRepository.DeleteClinicByID(id)
}
