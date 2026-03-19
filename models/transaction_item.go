package models

import "gorm.io/gorm"

type TransactionItem struct {
	gorm.Model
	TransactionID uint
	Transaction   Transaction
	ProductID     uint
	Product       Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product"`
	Quantity      int     `gorm:"not null" json:"quantity"`
	Price         float64 `gorm:"not null" json:"price"`
}
