package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DataBaseConnection() {
	godotenv.Load()
	
	var DSN = os.Getenv("DSN")
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)

	} else {
		log.Println("DB Connected")
	}

}
