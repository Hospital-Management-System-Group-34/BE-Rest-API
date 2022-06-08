package schedule

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
)

type getScheduleByIDUseCase struct {
	scheduleRepository domain.ScheduleRepository
}

func NewGetScheduleByIDUseCase(scheduleRepository domain.ScheduleRepository) domain.GetScheduleByIDUseCase {
	return &getScheduleByIDUseCase{
		scheduleRepository: scheduleRepository,
	}
}

func (u *getScheduleByIDUseCase) Execute(id uint) (entity.Schedule, int, error) {
	return u.scheduleRepository.GetScheduleByID(id)
}
