package services

import (
	"Synconomics/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAuditData(t *testing.T) {
	business := &models.Business{Name: "Test Shop"}
	transactions := []models.Transaction{
		{
			Status:      models.StatusCompleted,
			TotalAmount: 1000.0,
			TransactionItems: []models.TransactionItem{
				{Product: models.Product{Name: "Product A"}, Price: 100.0, Quantity: 10},
			},
		},
	}
	expenses := []models.Expense{
		{Category: "Rent", Amount: 200.0},
		{Category: "Utilities", Amount: 50.0},
	}
	products := []models.Product{
		{Name: "Product A", Price: 100, Stock: 50, MinStock: 10},
		{Name: "Product B", Price: 200, Stock: 5, MinStock: 10},
	}

	data := CalculateAuditData(business, transactions, expenses, products)

	assert.Equal(t, "Test Shop", data.BusinessName)
	assert.Equal(t, 1000.0, data.Metrics.Revenue)
	assert.Equal(t, 250.0, data.Metrics.TotalExpenses)
	assert.Equal(t, 750.0, data.Metrics.NetProfit)
	assert.Equal(t, 75.0, data.Metrics.NetProfitMargin)
	assert.Equal(t, 1, data.Metrics.LowStockCount) // Product B is low stock
	assert.Equal(t, 6000.0, data.Metrics.InventoryValue) // (100*50) + (200*5) = 5000 + 1000 = 6000
}
