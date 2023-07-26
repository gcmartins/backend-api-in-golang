package database

import (
	"MileTravel/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConfig() {
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Panic("Database connection error.")
	}

	DB.AutoMigrate(&models.Testimonial{})
}
