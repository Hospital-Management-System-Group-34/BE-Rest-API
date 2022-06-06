package doctor

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type addDoctorUseCase struct {
	doctorRepository domain.DoctorRepository
}

func NewAddDoctorUseCase(doctorRepository domain.DoctorRepository) domain.AddDoctorUseCase {
	return &addDoctorUseCase{
		doctorRepository: doctorRepository,
	}
}

func (u *addDoctorUseCase) Execute(payload entity.Doctor) (int, error) {
	return u.doctorRepository.AddDoctor(payload)
}
