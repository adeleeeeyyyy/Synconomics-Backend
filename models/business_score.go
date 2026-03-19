package models

import "gorm.io/gorm"

type BusinessScore struct {
	gorm.Model
	BusinessID uint   `json:"business_id"`

	Score      int    `json:"score" gorm:"type:int;not null"`
	Insight    string `json:"insight" gorm:"type:text"`
}