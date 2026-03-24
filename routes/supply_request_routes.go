package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupSupplyRequestRoutes(api fiber.Router) {
	repo := repositories.NewSupplyRequestRepository(config.DB)
	service := services.NewSupplyRequestService(repo)
	handler := handlers.NewSupplyRequestHandler(service)

	sr := api.Group("/supply-requests")

	sr.Use(handlers.JWTMiddleware)

	sr.Post("/", handler.CreateSupplyRequest)
	sr.Get("/", handler.GetSupplyRequests)
	sr.Get("/:id", handler.GetSupplyRequest)
	sr.Get("/business/:id", handler.GetSupplyRequestsByBusinessId)
	sr.Put("/:id", handler.UpdateSupplyRequest)
	sr.Delete("/:id", handler.DeleteSupplyRequest)
}
