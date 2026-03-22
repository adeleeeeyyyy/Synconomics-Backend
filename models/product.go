package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	BusinessID   uint `json:"business_id"`
	Business     Business
	SupplyOffers []SupplyOffer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:varchar(255)"`
	Price       int    `json:"price" gorm:"type:int;not null"`
	Stock       int    `json:"stock" gorm:"type:int;not null"`
	MinStock    int    `json:"min_stock" gorm:"type:int;not null"`
	ImageURL    string `json:"image_url" gorm:"type:varchar(255)"`
}
