package models

import "gorm.io/gorm"

type AIResult struct {
	gorm.Model
	AISessionID	uint
	AISession	AISession

	ResultType string    `gorm:"type:varchar(255)" json:"result_type"`
	Data       string    `gorm:"type:json" json:"data"`
}
