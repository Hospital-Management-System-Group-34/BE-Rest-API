package clinic

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type updateClinicByIDUseCase struct {
	clinicRepository domain.ClinicRepository
}

func NewUpdateClinicByIDUseCase(clinicRepository domain.ClinicRepository) domain.UpdateClinicByIDUseCase {
	return &updateClinicByIDUseCase{
		clinicRepository: clinicRepository,
	}
}

func (u *updateClinicByIDUseCase) Execute(payload entity.UpdateClinicPayload) (int, error) {
	return u.clinicRepository.UpdateClinicByID(payload)
}
