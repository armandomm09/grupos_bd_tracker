package handlers

import (
	"brithdayApi/models"
	"brithdayApi/database"
	"github.com/gofiber/fiber/v2"
)

func CreatePerson(c *fiber.Ctx) error {
	person := new(models.Person)

	if err := c.BodyParser(person); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if person.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email cannot be empty",
		})
	}

	result := database.DB.Create(&person)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : "Could not create person",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(person)
}

// func GetPeople(c *fiber.Ctx) error {
// 	var people []models.Person

// 	result := database.DB.Find(&people)
// }