package schedule

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type addScheduleUseCase struct {
	scheduleRepository domain.ScheduleRepository
	doctorRepository   domain.DoctorRepository
}

func NewAddScheduleUseCase(
	scheduleRepository domain.ScheduleRepository,
	doctorRepository domain.DoctorRepository,
) domain.AddScheduleUseCase {
	return &addScheduleUseCase{
		scheduleRepository: scheduleRepository,
		doctorRepository:   doctorRepository,
	}
}

func (u *addScheduleUseCase) Execute(payload entity.Schedule) (int, error) {
	if _, code, err := u.doctorRepository.GetDoctorByID(payload.DoctorID); err != nil {
		return code, err
	}

	return u.scheduleRepository.AddSchedule(payload)
}
