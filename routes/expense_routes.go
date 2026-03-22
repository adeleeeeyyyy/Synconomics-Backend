package routes

import (
	"Synconomics/config"
	"Synconomics/handlers"
	"Synconomics/repositories"
	"Synconomics/services"

	"github.com/gofiber/fiber/v2"
)

func SetupExpenseRoutes(api fiber.Router) {
	expenseRepo := repositories.NewExpenseRepository(config.DB)
	expenseService := services.NewExpenseService(expenseRepo)
	expenseHandler := handlers.NewExpenseHandler(expenseService)

	expenses := api.Group("/expenses")

	expenses.Use(handlers.JWTMiddleware)

	expenses.Post("/", expenseHandler.CreateExpense)
	expenses.Get("/", expenseHandler.GetExpenses)
	expenses.Get("/:id", expenseHandler.GetExpense)
	expenses.Get("/business/:id", expenseHandler.GetExpensesByBusinessId)
	expenses.Put("/:id", expenseHandler.UpdateExpense)
	expenses.Delete("/:id", expenseHandler.DeleteExpense)
}
