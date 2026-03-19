package models

import (
	"time"

	"gorm.io/gorm"
)

type BusinessMetric struct {
	gorm.Model

	BusinessID uint     `json:"business_id"`
	Business   Business `json:"business"`

	TotalRevenue  float64   `json:"total_revenue"`
	TotalExpense  float64   `json:"total_expense"`
	NetProfit     float64   `json:"net_profit"`
	StockTurnover float64   `json:"stock_turnover"`
	SalesGrowth   float64   `json:"sales_growth"`
	CalculatedAt  time.Time `json:"calculated_at"`
}
