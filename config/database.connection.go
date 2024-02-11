package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=root password=Seguridad2024 dbname=gorm port=5432"
var DB *gorm.DB

func DataBaseConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)

	} else {
		log.Println("DB Connected")
	}

}
