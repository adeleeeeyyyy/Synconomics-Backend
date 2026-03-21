package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/gofiber/swagger"

	"Synconomics/config"
	"Synconomics/routes"
	_ "Synconomics/docs"
)

// @title Synconomics API
// @version 1.0
// @description API for Synconomics, build with go
// @termsOfService http://swagger.io/terms/
// @host api-synconomics.synchronizeteams.com
// @BasePath /api
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer " followed by your JSON Web Token. Example: "Bearer eyJhb..."
func main() {
	godotenv.Load()

	// setup OAuth providers
	config.SetupOauth()

	// connect database
	config.ConnectDB()

	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, https://api-synconomics.synchronizeteams.com",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api")

	routes.SetupAuthRoutes(api)
	routes.SetupProductRoutes(api)
	routes.SetupBusinessRoutes(api)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
