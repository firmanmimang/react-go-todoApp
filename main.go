package main

import (
	"log"
	"os"

	"github.com/firmanmimang/react-go-todo/config"
	"github.com/firmanmimang/react-go-todo/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file if not in production
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	// Initialize the database connection
	config.ConnectDB()

	app := fiber.New()

	if os.Getenv("ENV") != "production" {
		app.Use(cors.New(cors.Config{
			AllowOrigins: "http://localhost:5173",
			AllowMethods: "GET,POST,PUT,DELETE,PATCH,OPTIONS",
			AllowHeaders: "Origin,Content-Type,Accept",
		}))
	}

	// Set up routes
	routes.Setup(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Serve static files in production
	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
