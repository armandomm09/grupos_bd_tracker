package main

import (
	"brithdayApi/routes"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func requestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Obtener la dirección IP del cliente
		clientIP := c.IP()

		// Imprimir detalles de la solicitud
		fmt.Println("\n- - - - - - - INCOMING REQUEST - - - - - - - -\n")
		fmt.Printf("Received request: %s %s\n", c.Method(), c.Path())
		fmt.Printf("Client IP: %s\n", clientIP)
		fmt.Printf("User Agent: %s\n", c.Get("User-Agent"))

		// Llamar al siguiente middleware o handler
		err := c.Next()

		fmt.Printf("Completed request: %s %s in %v\n", c.Method(), c.Path(), time.Since(start))

		return err
	}
}

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

	// database.ConnectDB()
	app.Use(requestLogger())
	routes.SetPersonRoutes(app)

	log.Fatal(app.Listen(":8888"))
}
