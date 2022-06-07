package clinic

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type getClinicByIDUseCase struct {
	clinicRepository domain.ClinicRepository
}

func NewGetClinicByIDUseCase(clinicRepository domain.ClinicRepository) domain.GetClinicByIDUseCase {
	return &getClinicByIDUseCase{
		clinicRepository: clinicRepository,
	}
}

func (u *getClinicByIDUseCase) Execute(id uint) (entity.Clinic, int, error) {
	return u.clinicRepository.GetClinicByID(id)
}
