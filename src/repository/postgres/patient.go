package postgres

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"gorm.io/gorm"
)

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) domain.PatientRepository {
	newPatientRepository := patientRepository{
		db: db,
	}

	return &newPatientRepository
}

func (r *patientRepository) AddPatient(payload entity.Patient) (int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *patientRepository) GetPatients() ([]entity.Patient, int, error) {
	Patients := []entity.Patient{}
	r.db.Find(&Patients)

	return Patients, http.StatusOK, nil
}

func (r *patientRepository) GetPatientByID(id uint) (entity.Patient, int, error) {
	Patient := entity.Patient{}
	result := r.db.Where("id = ?", id).First(&Patient)

	if result.RowsAffected == 0 {
		return entity.Patient{}, http.StatusNotFound, fmt.Errorf("Patient not found")
	}

	return Patient, http.StatusOK, nil
}

func (r *patientRepository) UpdatePatientByID(payload entity.UpdatePatientPayload) (int, error) {
	Patient, code, err := r.GetPatientByID(payload.ID)
	if err != nil {
		return code, err
	}

	Patient.FName = payload.FName
	Patient.LName = payload.LName
	Patient.Phone = payload.Phone
	Patient.Keluhan = payload.Keluhan

	result := r.db.Save(&Patient)

	if result.Error != nil {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *patientRepository) DeletePatientByID(id uint) (int, error) {
	result := r.db.Where("id = ?", id).Delete(&entity.Patient{})

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("Patient not found")
	}

	return http.StatusOK, nil
}
