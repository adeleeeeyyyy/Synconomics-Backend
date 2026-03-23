package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(api fiber.Router) {
	businessRepo := repositories.NewBusinessRepository(config.DB)
	businessService := services.NewBusinessService(businessRepo)

	productRepo := repositories.NewProductRepository(config.DB)
	productService := services.NewProductService(productRepo)

	productHandler := handlers.NewProductHandler(productService, businessService)

	products := api.Group("/products")

	products.Use(handlers.JWTMiddleware)

	products.Post("/", productHandler.CreateProduct)
	products.Get("/", productHandler.GetProducts)
	products.Get("/business/:businessId", productHandler.GetProductsByBusiness)
	products.Get("/:id", productHandler.GetProduct)
	products.Put("/:id", productHandler.UpdateProduct)
	products.Delete("/:id", productHandler.DeleteProduct)
}