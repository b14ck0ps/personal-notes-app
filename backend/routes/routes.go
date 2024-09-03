package routes

import (
	"github.com/b14ck0ps/personal-notes/controllers"
	"github.com/b14ck0ps/personal-notes/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Auth routes
	auth := app.Group("/auth")
	auth.Post("/register", controllers.RegisterUser)
	auth.Post("/login", controllers.LoginUser)

	// Note routes
	note := app.Group("/notes", middleware.JWTProtected())
	note.Get("/", controllers.GetNotes)
	note.Post("/", controllers.CreateNote)
	note.Get("/:id", controllers.GetNoteByID)
	note.Put("/:id", controllers.UpdateNote)
	note.Delete("/:id", controllers.DeleteNote)
}
