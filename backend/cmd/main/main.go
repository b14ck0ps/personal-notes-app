package main

import (
	"log"

	"github.com/b14ck0ps/personal-notes/database"
	"github.com/b14ck0ps/personal-notes/models"
	"github.com/b14ck0ps/personal-notes/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	database.DB.AutoMigrate(&models.User{}, &models.Note{})

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
