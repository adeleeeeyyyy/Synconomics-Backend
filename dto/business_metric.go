package dto

import (
	"Synconomics/models"
	"time"
)

type CreateBusinessMetricRequest struct {
	BusinessID    uint      `json:"business_id" validate:"required"`
	TotalRevenue  float64   `json:"total_revenue"`
	TotalExpense  float64   `json:"total_expense"`
	NetProfit     float64   `json:"net_profit"`
	StockTurnover float64   `json:"stock_turnover"`
	SalesGrowth   float64   `json:"sales_growth"`
	CalculatedAt  time.Time `json:"calculated_at"`
}

type UpdateBusinessMetricRequest struct {
	TotalRevenue  float64   `json:"total_revenue"`
	TotalExpense  float64   `json:"total_expense"`
	NetProfit     float64   `json:"net_profit"`
	StockTurnover float64   `json:"stock_turnover"`
	SalesGrowth   float64   `json:"sales_growth"`
	CalculatedAt  time.Time `json:"calculated_at"`
}

type BusinessMetricResponse struct {
	ID            uint            `json:"id"`
	BusinessID    uint            `json:"business_id"`
	Business      models.Business `json:"business"`
	TotalRevenue  float64         `json:"total_revenue"`
	TotalExpense  float64         `json:"total_expense"`
	NetProfit     float64         `json:"net_profit"`
	StockTurnover float64         `json:"stock_turnover"`
	SalesGrowth   float64         `json:"sales_growth"`
	CalculatedAt  time.Time       `json:"calculated_at"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}
