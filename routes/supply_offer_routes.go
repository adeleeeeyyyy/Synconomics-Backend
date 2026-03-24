package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupSupplyOfferRoutes(api fiber.Router) {
	repo := repositories.NewSupplyOfferRepository(config.DB)
	service := services.NewSupplyOfferService(repo)
	handler := handlers.NewSupplyOfferHandler(service)

	so := api.Group("/supply-offers")

	so.Use(handlers.JWTMiddleware)

	so.Post("/", handler.CreateSupplyOffer)
	so.Get("/", handler.GetSupplyOffers)
	so.Get("/:id", handler.GetSupplyOffer)
	so.Get("/business/:id", handler.GetSupplyOffersByBusinessId)
	so.Put("/:id", handler.UpdateSupplyOffer)
	so.Delete("/:id", handler.DeleteSupplyOffer)
}
