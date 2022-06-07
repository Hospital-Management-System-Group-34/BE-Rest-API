package patient

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type addPatientUseCase struct {
	patientRepository domain.PatientRepository
}

func NewAddPatientUseCase(patientRepository domain.PatientRepository) domain.AddPatientUseCase {
	return &addPatientUseCase{
		patientRepository: patientRepository,
	}
}

func (u *addPatientUseCase) Execute(payload entity.Patient) (int, error) {
	return u.patientRepository.AddPatient(payload)
}
