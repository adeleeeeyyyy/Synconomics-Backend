package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupBusinessMetricRoutes(api fiber.Router) {
	repo := repositories.NewBusinessMetricRepository(config.DB)
	service := services.NewBusinessMetricService(repo)
	handler := handlers.NewBusinessMetricHandler(service)

	bm := api.Group("/business-metrics")
	bm.Use(handlers.JWTMiddleware)

	bm.Post("/", handler.CreateMetric)
	bm.Get("/", handler.GetMetrics)
	bm.Get("/:id", handler.GetMetric)
	bm.Get("/business/:id", handler.GetMetricsByBusinessId)
	bm.Get("/business/:id/latest", handler.GetLatestMetricByBusinessId)
	bm.Put("/:id", handler.UpdateMetric)
	bm.Delete("/:id", handler.DeleteMetric)
}
