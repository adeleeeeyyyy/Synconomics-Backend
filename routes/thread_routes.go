package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupThreadRoutes(api fiber.Router) {
	repo := repositories.NewThreadRepository(config.DB)
	service := services.NewThreadService(repo)
	handler := handlers.NewThreadHandler(service)

	th := api.Group("/threads")

	th.Use(handlers.JWTMiddleware)

	th.Post("/", handler.CreateThread)
	th.Get("/", handler.GetThreads)
	th.Get("/:id", handler.GetThread)
	th.Put("/:id", handler.UpdateThread)
	th.Delete("/:id", handler.DeleteThread)
}
