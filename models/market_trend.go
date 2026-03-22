package models

import (
	"gorm.io/gorm"
)

type MarketTrend struct {
	gorm.Model
	ProductName string         `gorm:"type:varchar(255);not null" json:"product_name"`
	SearchCount int            `gorm:"not null" json:"search_count"`
	DemandScore float64        `gorm:"not null" json:"demand_score"`
	Location    string         `gorm:"type:varchar(255);not null" json:"location"`
}
