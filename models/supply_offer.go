package models

import "gorm.io/gorm"

type SupplyOffer struct {
	gorm.Model
	BusinessID uint     `json:"business_id"`
	Business   Business `json:"business"`

	ProductID uint    `json:"product_id"`
	Product   Product `json:"product"`

	SupplyMatches	[]SupplyMatch `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ProductName string `json:"product_name" gorm:"size:255;not null"`
	Quantity    int    `json:"quantity" gorm:"not null"`
}
