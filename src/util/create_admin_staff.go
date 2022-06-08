package util

import (
	"os"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"golang.org/x/crypto/bcrypt"
)

func CreateAdminStaff() {
	db := postgres.Connect()

	staff := entity.Staff{}
	result := db.Where("staff_type = ?", "admin").First(&staff)

	if result.RowsAffected == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)
		db.Create(&entity.Staff{
			StaffType: "admin",
			Name:      "Admin Simars",
			Phone:     os.Getenv("ADMIN_PHONE"),
			Email:     os.Getenv("ADMIN_EMAIL"),
			Password:  string(hashedPassword),
			Token:     "token",
		})
	}
}
