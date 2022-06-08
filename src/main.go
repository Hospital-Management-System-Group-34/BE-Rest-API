package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/server"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/util"

	"github.com/joho/godotenv"
)

func init() {
	environment := os.Getenv("ENVIRONMENT")

	if environment != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error load .env file: ", err.Error())
		}
	}

	postgres.InitMigration()
	util.CreateAdminStaff()
}

func main() {
	e := server.CreateServer()

	if err := e.Start(os.Getenv("PORT")); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
