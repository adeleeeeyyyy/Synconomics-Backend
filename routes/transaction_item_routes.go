package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupTransactionItemRoutes(api fiber.Router) {
	transactionItemRepo := repositories.NewTransactionItemRepository(config.DB)
	productRepo := repositories.NewProductRepository(config.DB)
	transactionItemService := services.NewTransactionItemService(transactionItemRepo, productRepo)
	transactionItemHandler := handlers.NewTransactionItemHandler(transactionItemService)

	transactionItems := api.Group("/transaction_items")

	transactionItems.Use(handlers.JWTMiddleware)

	transactionItems.Post("/", transactionItemHandler.CreateTransactionItem)
	transactionItems.Get("/", transactionItemHandler.GetTransactionItems)
	transactionItems.Get("/:id", transactionItemHandler.GetTransactionItem)
	transactionItems.Get("/transaction/:id", transactionItemHandler.GetTransactionItemsByTransactionId)
	transactionItems.Put("/:id", transactionItemHandler.UpdateTransactionItem)
	transactionItems.Delete("/:id", transactionItemHandler.DeleteTransactionItem)
}
