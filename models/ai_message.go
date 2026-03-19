package models

import "gorm.io/gorm"

type AIMessage struct {
	gorm.Model
	AISessionID		uint
	AISession		AISession
	Role    string `gorm:"type:varchar(50);not null" json:"role"`
	Content string `gorm:"type:text;not null" json:"content"`
}
