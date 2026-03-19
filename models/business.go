package models

import "gorm.io/gorm"

type Business struct {
	gorm.Model
	UserID          uint
	User            User             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Foreign Key to User
	Products        []Product        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Transactions    []Transaction    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Expenses        []Expense        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AISessions      []AISession      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SupplyRequests  []SupplyRequest  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SupplyOffers    []SupplyOffer    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BusinessMetrics []BusinessMetric `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BusinessScores  []BusinessScore  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Name        string  `json:"name" gorm:"type:varchar(100)"`
	Description string  `json:"description" gorm:"type:varchar(255)"`
	Category    string  `json:"category" gorm:"type:varchar(100)"`
	LogoURL     string  `json:"logo_url" gorm:"type:varchar(255)"`
	Address     string  `json:"address" gorm:"type:varchar(255)"`
	Latitude    float32 `json:"latitude" gorm:"type:float"`
	Longitude   float32 `json:"longitude" gorm:"type:float"`
	Phone       string  `json:"phone" gorm:"type:varchar(255)"`
	Whatsapp    string  `json:"whatsapp" gorm:"type:varchar(255)"`
	Instagram   string  `json:"instagram" gorm:"type:varchar(255)"`
	Tiktok      string  `json:"tiktok" gorm:"type:varchar(255)"`
	Website     string  `json:"website" gorm:"type:varchar(255)"`
}