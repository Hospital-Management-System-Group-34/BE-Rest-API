package schedule

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type updateScheduleByIDUseCase struct {
	scheduleRepository domain.ScheduleRepository
	doctorRepository   domain.DoctorRepository
}

func NewUpdateScheduleByIDUseCase(
	scheduleRepository domain.ScheduleRepository,
	doctorRepository domain.DoctorRepository,
) domain.UpdateScheduleByIDUseCase {
	return &updateScheduleByIDUseCase{
		scheduleRepository: scheduleRepository,
		doctorRepository:   doctorRepository,
	}
}

func (u *updateScheduleByIDUseCase) Execute(payload entity.UpdateSchedulePayload) (int, error) {
	if _, code, err := u.doctorRepository.GetDoctorByID(payload.DoctorID); err != nil {
		return code, err
	}

	return u.scheduleRepository.UpdateScheduleByID(payload)
}
