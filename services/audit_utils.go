package services

import (
	"Synconomics/models"
)

type AuditMetrics struct {
	Revenue           float64 `json:"revenue"`
	TotalExpenses     float64 `json:"total_expenses"`
	GrossProfit       float64 `json:"gross_profit"`
	NetProfit         float64 `json:"net_profit"`
	GrossProfitMargin float64 `json:"gross_profit_margin"`
	NetProfitMargin   float64 `json:"net_profit_margin"`
	InventoryValue    float64 `json:"inventory_value"`
	LowStockCount     int     `json:"low_stock_count"`
	TransactionCount  int     `json:"transaction_count"`
}

type CategorySummary struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type AuditData struct {
	BusinessName    string            `json:"business_name"`
	Period          string            `json:"period"`
	Metrics         AuditMetrics      `json:"metrics"`
	ExpenseAnalysis []CategorySummary `json:"expense_analysis"`
	TopProducts     []CategorySummary `json:"top_products"`
}

func CalculateAuditData(
	business *models.Business,
	transactions []models.Transaction,
	expenses []models.Expense,
	products []models.Product,
) AuditData {
	data := AuditData{
		BusinessName: business.Name,
		Period:       "Last 30 Days",
	}

	// 1. Calculate Revenue and Transaction Stats
	revenue := 0.0
	productSales := make(map[string]float64)
	for _, tx := range transactions {
		if tx.Status == models.StatusCompleted {
			revenue += tx.TotalAmount
			for _, item := range tx.TransactionItems {
				productSales[item.Product.Name] += item.Price * float64(item.Quantity)
			}
		}
	}
	data.Metrics.Revenue = revenue
	data.Metrics.TransactionCount = len(transactions)

	// 2. Calculate Expenses
	totalExpenses := 0.0
	expenseCats := make(map[string]float64)
	for _, e := range expenses {
		totalExpenses += e.Amount
		expenseCats[e.Category] += e.Amount
	}
	data.Metrics.TotalExpenses = totalExpenses

	// 3. Profitability
	data.Metrics.GrossProfit = revenue // COGS not easily available without manual cost field, using simple revenue as proxy for now
	data.Metrics.NetProfit = revenue - totalExpenses
	if revenue > 0 {
		data.Metrics.GrossProfitMargin = (data.Metrics.GrossProfit / revenue) * 100
		data.Metrics.NetProfitMargin = (data.Metrics.NetProfit / revenue) * 100
	}

	// 4. Inventory
	inventoryValue := 0.0
	lowStock := 0
	for _, p := range products {
		inventoryValue += float64(p.Price * p.Stock)
		if p.Stock <= p.MinStock {
			lowStock++
		}
	}
	data.Metrics.InventoryValue = inventoryValue
	data.Metrics.LowStockCount = lowStock

	// 5. Aggregate Summaries
	for cat, amt := range expenseCats {
		data.ExpenseAnalysis = append(data.ExpenseAnalysis, CategorySummary{Name: cat, Amount: amt})
	}
	for prod, amt := range productSales {
		data.TopProducts = append(data.TopProducts, CategorySummary{Name: prod, Amount: amt})
	}

	return data
}
