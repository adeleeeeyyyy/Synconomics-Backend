package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupBusinessRoutes(api fiber.Router) {
	businessRepo := repositories.NewBusinessRepository(config.DB)
	businessService := services.NewBusinessService(businessRepo)
	businessHandler := handlers.NewBusinessHandler(businessService)

	business := api.Group("/business")

	// Pasang JWT Middleware agar semua endpoint business butuh Authorization Bearer
	business.Use(handlers.JWTMiddleware)

	business.Post("/", businessHandler.CreateBusiness)
	business.Get("/", businessHandler.GetAllBusinesses)
	business.Get("/:id", businessHandler.GetBusinessById)
	business.Put("/:id", businessHandler.UpdateBusiness)
	business.Delete("/:id", businessHandler.DeleteBusiness)
}
