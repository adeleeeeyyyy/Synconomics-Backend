package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupSupplyMatchRoutes(api fiber.Router) {
	repo := repositories.NewSupplyMatchRepository(config.DB)
	service := services.NewSupplyMatchService(repo)
	handler := handlers.NewSupplyMatchHandler(service)

	sm := api.Group("/supply-matches")

	sm.Use(handlers.JWTMiddleware)

	sm.Post("/", handler.CreateSupplyMatch)
	sm.Get("/", handler.GetSupplyMatches)
	sm.Get("/:id", handler.GetSupplyMatch)
	sm.Patch("/:id/status", handler.UpdateSupplyMatchStatus)
	sm.Delete("/:id", handler.DeleteSupplyMatch)
}
