package doctor

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type updateDoctorByIDUseCase struct {
	doctorRepository domain.DoctorRepository
}

func NewUpdateDoctorByIDUseCase(doctorRepository domain.DoctorRepository) domain.UpdateDoctorByIDUseCase {
	return &updateDoctorByIDUseCase{
		doctorRepository: doctorRepository,
	}
}

func (u *updateDoctorByIDUseCase) Execute(payload entity.UpdateDoctorPayload) (int, error) {
	return u.doctorRepository.UpdateDoctorByID(payload)
}
