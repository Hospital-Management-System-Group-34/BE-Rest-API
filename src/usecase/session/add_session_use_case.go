package session

import (
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type addSessionUseCase struct {
	sessionRepository  domain.SessionRepository
	patientRepository  domain.PatientRepository
	clinicRepository   domain.ClinicRepository
	doctorRepository   domain.DoctorRepository
	scheduleRepository domain.ScheduleRepository
}

func NewAddSessionUseCase(
	sessionRepository domain.SessionRepository,
	patientRepository domain.PatientRepository,
	clinicRepository domain.ClinicRepository,
	doctorRepository domain.DoctorRepository,
	scheduleRepository domain.ScheduleRepository,
) domain.AddSessionUseCase {
	return &addSessionUseCase{
		sessionRepository:  sessionRepository,
		patientRepository:  patientRepository,
		clinicRepository:   clinicRepository,
		doctorRepository:   doctorRepository,
		scheduleRepository: scheduleRepository,
	}
}

func (u *addSessionUseCase) Execute(payload entity.Session) (entity.Session, int, error) {
	// panggil get by id patient, clinic, doctor, schedule, buat verifikasi
	if _, code, err := u.patientRepository.GetPatientByID(payload.PatientID); err != nil {
		return entity.Session{}, code, err
	}

	// if _, code, err := u.clinicRepository.GetClinicByID(payload.ClinicID); err != nil {
	// 	return entity.Session{}, code, err
	// }

	if _, code, err := u.doctorRepository.GetDoctorByID(payload.DoctorID); err != nil {
		return entity.Session{}, code, err
	}

	if _, code, err := u.scheduleRepository.GetScheduleByID(payload.ScheduleID); err != nil {
		return entity.Session{}, code, err
	}

	queue, code, err := u.sessionRepository.GetSessionLastQueue(payload.ScheduleID)
	if err != nil {
		return entity.Session{}, code, err
	}

	payload.Queue = queue + 1

	session, code, err := u.sessionRepository.AddSession(payload)
	if err != nil {
		return entity.Session{}, code, err
	}

	return session, http.StatusOK, nil
}
