package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupMarketTrendRoutes(api fiber.Router) {
	repo := repositories.NewMarketTrendRepository(config.DB)
	aiRepo := repositories.NewAIRepository(config.DB)
	logRepo := repositories.NewProductSearchLogRepository(config.DB)
	
	aiService := services.NewAIService(aiRepo)
	service := services.NewMarketTrendService(repo, aiService, logRepo)
	handler := handlers.NewMarketTrendHandler(service)

	trends := api.Group("/market-trends")
	trends.Use(handlers.JWTMiddleware)

	trends.Post("/", handler.CreateTrend)
	trends.Get("/", handler.GetAllTrends)
	trends.Get("/top", handler.GetTopTrends)
	trends.Post("/refresh", handler.RefreshTrends)
	trends.Get("/:id", handler.GetTrendById)
	trends.Put("/:id", handler.UpdateTrend)
	trends.Delete("/:id", handler.DeleteTrend)
}
