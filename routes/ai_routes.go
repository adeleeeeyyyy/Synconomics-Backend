package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupAIRoutes(api fiber.Router) {
	aiRepo := repositories.NewAIRepository(config.DB)
	aiService := services.NewAIService(aiRepo)
	aiHandler := handlers.NewAIHandler(aiService)

	aiRoutes := api.Group("/ai")

	// Protected routes
	aiRoutes.Use(handlers.JWTMiddleware)

	aiRoutes.Post("/sessions", aiHandler.CreateSession)
	aiRoutes.Post("/sessions/:id/chat", aiHandler.Chat)
	aiRoutes.Get("/sessions/:id/messages", aiHandler.GetMessages)
	aiRoutes.Post("/sessions/:id/result", aiHandler.FinalizeResult)
}
