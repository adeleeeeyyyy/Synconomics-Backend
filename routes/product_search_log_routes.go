package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupProductSearchLogRoutes(api fiber.Router) {
	repo := repositories.NewProductSearchLogRepository(config.DB)
	service := services.NewProductSearchLogService(repo)
	handler := handlers.NewProductSearchLogHandler(service)

	psl := api.Group("/product-search-logs")

	psl.Use(handlers.JWTMiddleware)

	psl.Post("/", handler.CreateLog)
	psl.Get("/", handler.GetLogs)
	psl.Get("/me", handler.GetUserLogs)
	psl.Get("/:id", handler.GetLog)
}
