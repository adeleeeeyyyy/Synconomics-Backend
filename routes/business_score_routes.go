package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupBusinessScoreRoutes(api fiber.Router) {
	repo := repositories.NewBusinessScoreRepository(config.DB)
	service := services.NewBusinessScoreService(repo)
	handler := handlers.NewBusinessScoreHandler(service)

	bs := api.Group("/business-scores")
	bs.Use(handlers.JWTMiddleware)

	bs.Post("/", handler.CreateScore)
	bs.Get("/", handler.GetScores)
	bs.Get("/:id", handler.GetScore)
	bs.Get("/business/:id", handler.GetScoresByBusinessId)
	bs.Get("/business/:id/latest", handler.GetLatestScoreByBusinessId)
	bs.Put("/:id", handler.UpdateScore)
	bs.Delete("/:id", handler.DeleteScore)
}
