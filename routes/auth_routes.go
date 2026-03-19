package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(api fiber.Router) {
	userRepo := repositories.NewUserRepository(config.DB)
	authService := services.NewAuthServices(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	api.Get("/auth/google", authHandler.GoogleLogin)
	api.Get("/auth/google/callback", authHandler.GoogleCallback)

	api.Get("/profile", handlers.JWTMiddleware, authHandler.Profile)
}