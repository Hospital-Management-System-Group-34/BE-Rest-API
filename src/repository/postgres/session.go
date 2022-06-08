package postgres

import (
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"gorm.io/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) domain.SessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (r *sessionRepository) AddSession(payload entity.Session) (entity.Session, int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return entity.Session{}, http.StatusInternalServerError, result.Error
	}

	return payload, http.StatusOK, nil
}

func (r *sessionRepository) GetSessionLastQueue(scheduleID uint) (int, int, error) {
	lastQueue := 0

	result := r.db.Find(&entity.Session{})
	if result.RowsAffected != 0 {
		row := r.db.Table("sessions").Where("schedule_id = ?", scheduleID).Select("MAX(queue)").Row()
		err := row.Scan(&lastQueue)

		if err != nil {
			return -1, http.StatusInternalServerError, err
		}
	}

	if result.Error != nil {
		return -1, http.StatusInternalServerError, result.Error
	}

	return lastQueue, http.StatusOK, nil
}
