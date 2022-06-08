package schedule

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
)

type deleteScheduleByIDUseCase struct {
	scheduleRepository domain.ScheduleRepository
}

func NewDeleteScheduleByIDUseCase(scheduleRepository domain.ScheduleRepository) domain.DeleteScheduleByIDUseCase {
	return &deleteScheduleByIDUseCase{
		scheduleRepository: scheduleRepository,
	}
}

func (u *deleteScheduleByIDUseCase) Execute(id uint) (int, error) {
	return u.scheduleRepository.DeleteScheduleByID(id)
}
