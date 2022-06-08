package postgres

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"gorm.io/gorm"
)

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) domain.ScheduleRepository {
	return &scheduleRepository{
		db: db,
	}
}

func (r *scheduleRepository) AddSchedule(payload entity.Schedule) (int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *scheduleRepository) GetSchedules() ([]entity.Schedule, int, error) {
	schedules := []entity.Schedule{}
	r.db.Find(&schedules)

	return schedules, http.StatusOK, nil
}

func (r *scheduleRepository) GetScheduleByID(id uint) (entity.Schedule, int, error) {
	schedule := entity.Schedule{}
	result := r.db.Where("id = ?", id).First(&schedule)

	if result.RowsAffected == 0 {
		return entity.Schedule{}, http.StatusNotFound, fmt.Errorf("schedule not found")
	}

	return schedule, http.StatusOK, nil
}

func (r *scheduleRepository) UpdateScheduleByID(payload entity.UpdateSchedulePayload) (int, error) {
	schedule, code, err := r.GetScheduleByID(payload.ID)
	if err != nil {
		return code, err
	}

	schedule.Day = payload.Day
	schedule.Time = payload.Time
	schedule.DoctorID = payload.DoctorID

	result := r.db.Save(&schedule)

	if result.Error != nil {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *scheduleRepository) DeleteScheduleByID(id uint) (int, error) {
	result := r.db.Where("id = ?", id).Delete(&entity.Schedule{})

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("schedule not found")
	}

	return http.StatusOK, nil
}
