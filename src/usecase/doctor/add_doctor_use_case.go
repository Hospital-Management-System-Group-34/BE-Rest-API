package doctor

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type addDoctorUseCase struct {
	doctorRepository domain.DoctorRepository
	clinicRepository domain.ClinicRepository
}

func NewAddDoctorUseCase(
	doctorRepository domain.DoctorRepository,
	clinicRepository domain.ClinicRepository,
) domain.AddDoctorUseCase {
	return &addDoctorUseCase{
		doctorRepository: doctorRepository,
		clinicRepository: clinicRepository,
	}
}

func (u *addDoctorUseCase) Execute(payload entity.Doctor) (int, error) {
	if _, code, err := u.clinicRepository.GetClinicByID(payload.ClinicID); err != nil {
		return code, err
	}

	return u.doctorRepository.AddDoctor(payload)
}
