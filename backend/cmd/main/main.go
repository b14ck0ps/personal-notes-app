package main

import (
	"log"

	"github.com/b14ck0ps/personal-notes/models"
	"github.com/b14ck0ps/personal-notes/routes"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Initialize the app
	app := fiber.New()

	// PostgreSQL connection string
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdbname port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// Initialize the database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the database
	database.AutoMigrate(&models.User{}, &models.Note{})

	// Setup routes
	routes.SetupRoutes(app)

	// Start the app
	log.Fatal(app.Listen(":3000"))
}
