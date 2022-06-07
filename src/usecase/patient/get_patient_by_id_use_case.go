package patient

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type getPatientByIDUseCase struct {
	patientRepository domain.PatientRepository
}

func NewGetPatientByIDUseCase(patientRepository domain.PatientRepository) domain.GetPatientByIDUseCase {
	return &getPatientByIDUseCase{
		patientRepository: patientRepository,
	}
}

func (u *getPatientByIDUseCase) Execute(id uint) (entity.Patient, int, error) {
	return u.patientRepository.GetPatientByID(id)
}
