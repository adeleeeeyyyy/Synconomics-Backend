package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupReplyRoutes(api fiber.Router) {
	repo := repositories.NewReplyRepository(config.DB)
	service := services.NewReplyService(repo)
	handler := handlers.NewReplyHandler(service)

	re := api.Group("/replies")

	re.Use(handlers.JWTMiddleware)

	re.Post("/", handler.CreateReply)
	re.Get("/thread/:threadId", handler.GetRepliesByThread)
	re.Put("/:id", handler.UpdateReply)
	re.Delete("/:id", handler.DeleteReply)
}
