package models

import "gorm.io/gorm"

type SupplyRequestStatus string

const (
	StatusOpen    SupplyRequestStatus = "open"
	StatusMatched SupplyRequestStatus = "matched"
	StatusClosed  SupplyRequestStatus = "closed"
)

type SupplyRequest struct {
	gorm.Model
	BusinessID 	uint
	Business	Business
	SupplyMatches	[]SupplyMatch `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ProductName string              `json:"product_name"`
	Quantity    int                 `json:"quantity"`
	Status      SupplyRequestStatus `gorm:"type:enum('open','matched','closed');default:'open'" json:"status"`
}
