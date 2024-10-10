package main

import (
	"brithdayApi/database"
	"brithdayApi/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://0.0.0.0:80", // React front-end origin
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS, PATCH",
	}))

	database.ConnectDB()

	routes.SetPersonRoutes(app)

	log.Fatal(app.Listen(":4000"))
}