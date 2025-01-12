package db

import (
	"log"
	"umbrella/internal/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&models.User{},
		&models.MedicalDoctor{},
	)
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	log.Println("Database migration completed.")
}
