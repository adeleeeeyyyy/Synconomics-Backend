package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionStatus string

const (
	StatusPending   TransactionStatus = "pending"
	StatusCompleted TransactionStatus = "completed"
	StatusCancelled TransactionStatus = "cancelled"
)

type Transaction struct {
	gorm.Model
	TransactionItems []TransactionItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BusinessID       uint              `json:"business_id"`
	Business         Business          `json:"business"`

	TotalAmount     float64           `gorm:"type:decimal(10,2)" json:"total_amount"`
	PaymentMethod   string            `json:"payment_method"`
	Status          TransactionStatus `gorm:"type:enum('pending','completed','cancelled')" json:"status"`
	TransactionDate time.Time         `json:"transaction_date"`
}
