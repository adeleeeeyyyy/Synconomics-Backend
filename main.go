package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"
)

func main() {
	godotenv.Load()

	// setup OAuth providers
	config.SetupOauth()

	// connect database
	config.ConnectDB()

	// dependency injection
	userRepo := repositories.NewUserRepository(config.DB)
	authService := services.NewAuthServices(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	// public routes
	api := app.Group("/api")
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	// Google OAuth routes
	api.Get("/auth/google", authHandler.GoogleLogin)
	api.Get("/auth/google/callback", authHandler.GoogleCallback)

	// protected routes
	api.Get("/profile", handlers.JWTMiddleware, authHandler.Profile)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
