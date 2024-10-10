package routes

import (
	"github.com/gofiber/fiber/v2"
	"brithdayApi/handlers"
)

func SetPersonRoutes(app *fiber.App) {
	personRoutes := app.Group("/api/person")

	personRoutes.Post("/", handlers.CreatePerson)
}