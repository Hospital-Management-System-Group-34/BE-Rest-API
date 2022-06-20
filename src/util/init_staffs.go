package util

import (
	"fmt"
	"strings"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"golang.org/x/crypto/bcrypt"
)

func InitStaffs() {
	db := postgres.Connect()

	result := db.Where("role = ?", "Staff").Find(&entity.User{})

	if result.RowsAffected == 0 {
		staffs := []entity.User{
			{Name: "Jane Foster"},
			{Name: "Reed Richard"},
			{Name: "Peter Parker"},
			{Name: "Wanda Maximoff"},
			{Name: "Stephen Strange"},
			{Name: "Kamala Khan"},
		}

		for index, staff := range staffs {
			staff.ID = fmt.Sprintf("staff-%d", index+1)

			staff.Password = strings.ToLower(strings.Replace(staff.Name, " ", "", 1))
			hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(staff.Password), bcrypt.DefaultCost)
			staff.Password = string(hashedPassword)

			staff.Phone = "08123456789"
			staff.Role = "Staff"

			db.Create(&staff)
		}
	}
}
