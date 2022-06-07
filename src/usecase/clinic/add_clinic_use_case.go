package clinic

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type addClinicUseCase struct {
	clinicRepository domain.ClinicRepository
}

func NewAddClinicUseCase(clinicRepository domain.ClinicRepository) domain.AddClinicUseCase {
	return &addClinicUseCase{
		clinicRepository: clinicRepository,
	}
}

func (u *addClinicUseCase) Execute(payload entity.Clinic) (int, error) {
	return u.clinicRepository.AddClinic(payload)
}
