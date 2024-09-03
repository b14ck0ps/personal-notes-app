package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the database connection
func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=personal_notes port=5432 sslmode=disable TimeZone=Asia/Dhaka"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Assign DB connection to the global variable
	DB = database
}
