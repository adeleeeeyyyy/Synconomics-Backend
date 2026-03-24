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
	businessRepo := repositories.NewBusinessRepository(config.DB)

	authService := services.NewAuthServices(userRepo)
	businessService := services.NewBusinessService(businessRepo)

	authHandler := handlers.NewAuthHandler(authService, businessService)

	authGroup := api.Group("/auth")
	authGroup.Post("/register", authHandler.Register)
	authGroup.Post("/login", authHandler.Login)
	authGroup.Get("/google", authHandler.GoogleLogin)
	authGroup.Get("/google/callback", authHandler.GoogleCallback)

	// Protected routes
	authGroup.Use(handlers.JWTMiddleware)
	authGroup.Get("/profile", authHandler.Profile)
	authGroup.Put("/profile", authHandler.UpdateProfile)
	authGroup.Get("/me-with-businesses", authHandler.GetMeWithBusinesses)
}