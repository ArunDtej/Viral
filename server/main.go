package main

import (
	"log"
	"Viral/db"
	"Viral/handlers"
	// "time"

	"Viral/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	// "github.com/robfig/cron/v3"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	if err := db.InitDB(); err != nil {
		log.Fatal("Failed to connect to the database")
	}
	
	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ServerHeader: "Ubac",
		AppName:      "Ubac API v1.0",
	})

	app.Use(middlewares.SanitizeInputMiddleware)

	// Enable CORS for all origins and common headers
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, x-api-key",
	}))

	// Register modules
	app.Post("/api/page", handlers.HandlePage)
	app.Post("/api/predict", handlers.HandlePredict)
	app.Get("/api/prediction/:id", handlers.HandleGetPrediction)
	
	log.Fatal(app.Listen(":8000"))
}

