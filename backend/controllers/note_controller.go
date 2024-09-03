package controllers

import (
	"github.com/b14ck0ps/personal-notes/database"
	"github.com/b14ck0ps/personal-notes/models"
	"github.com/gofiber/fiber/v2"
)

// CreateNote creates a new note
func CreateNote(c *fiber.Ctx) error {
	var note models.Note
	if err := c.BodyParser(&note); err != nil {
		return err
	}

	database.DB.Create(&note)
	return c.JSON(note)
}

// GetNotes retrieves all notes for a user
func GetNotes(c *fiber.Ctx) error {
	var notes []models.Note
	userId := c.Locals("userId").(uint)

	database.DB.Where("user_id = ?", userId).Find(&notes)
	return c.JSON(notes)
}

// GetNoteByID retrieves a specific note by ID
func GetNoteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var note models.Note

	database.DB.First(&note, id)

	if note.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Note not found",
		})
	}

	return c.JSON(note)
}

// UpdateNote updates a note
func UpdateNote(c *fiber.Ctx) error {
	id := c.Params("id")
	var note models.Note

	database.DB.First(&note, id)

	if note.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Note not found",
		})
	}

	if err := c.BodyParser(&note); err != nil {
		return err
	}

	database.DB.Save(&note)
	return c.JSON(note)
}

// DeleteNote deletes a note
func DeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")
	var note models.Note

	database.DB.First(&note, id)

	if note.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Note not found",
		})
	}

	database.DB.Delete(&note)
	return c.SendStatus(204)
}
