package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupTransactionRoutes(api fiber.Router) {
	transactionRepo := repositories.NewTransactionRepository(config.DB)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	transactions := api.Group("/transactions")

	transactions.Use(handlers.JWTMiddleware)

	transactions.Post("/", transactionHandler.CreateTransaction)
	transactions.Get("/", transactionHandler.GetTransactions)
	transactions.Get("/:id", transactionHandler.GetTransaction)
	transactions.Get("/business/:id", transactionHandler.GetTransactionsByBusinessId)
	transactions.Put("/:id", transactionHandler.UpdateTransaction)
	transactions.Delete("/:id", transactionHandler.DeleteTransaction)
}
