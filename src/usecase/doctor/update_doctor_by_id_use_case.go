package doctor

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type updateDoctorByIDUseCase struct {
	doctorRepository domain.DoctorRepository
	clinicRepository domain.ClinicRepository
}

func NewUpdateDoctorByIDUseCase(
	doctorRepository domain.DoctorRepository,
	clinicRepository domain.ClinicRepository,
) domain.UpdateDoctorByIDUseCase {
	return &updateDoctorByIDUseCase{
		doctorRepository: doctorRepository,
		clinicRepository: clinicRepository,
	}
}

func (u *updateDoctorByIDUseCase) Execute(payload entity.UpdateDoctorPayload) (int, error) {
	if _, code, err := u.clinicRepository.GetClinicByID(payload.ClinicID); err != nil {
		return code, err
	}

	return u.doctorRepository.UpdateDoctorByID(payload)
}
