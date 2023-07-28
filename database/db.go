package database

import (
	"MileTravel/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConfig() {
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Panic("Database connection error.")
	}

	DB.AutoMigrate(&models.Testimonial{})
}

func TestDbConfig() {
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panic("Database connection error.")
	}

	DB.AutoMigrate(&models.Testimonial{})
}

func ClearTestDb() {
	e := os.Remove("test.db")
	if e != nil {
		log.Fatal(e)
	}
}
