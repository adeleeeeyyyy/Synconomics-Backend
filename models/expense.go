package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	BusinessID uint     
	Business   Business `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"business"`

	Title      string   `gorm:"type:varchar(255);not null" json:"title"`
	Amount     float64  `gorm:"type:decimal(10,2);not null" json:"amount"`
	Category   string   `gorm:"type:varchar(100);not null" json:"category"`
}
