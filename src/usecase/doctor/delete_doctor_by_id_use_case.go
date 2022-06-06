package doctor

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
)

type deleteDoctorByIDUseCase struct {
	doctorRepository domain.DoctorRepository
}

func NewDeleteDoctorByIDUseCase(doctorRepository domain.DoctorRepository) domain.DeleteDoctorByIDUseCase {
	return &deleteDoctorByIDUseCase{
		doctorRepository: doctorRepository,
	}
}

func (u *deleteDoctorByIDUseCase) Execute(id uint) (int, error) {
	return u.doctorRepository.DeleteDoctorByID(id)
}
