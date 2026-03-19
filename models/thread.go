package models

import "gorm.io/gorm"

type Thread struct {
	gorm.Model
	UserID uint `json:"user_id"`
	User   User `json:"user"`
	Replies []Reply `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Title   string `json:"title" gorm:"type:varchar(255);not null"`
	Content string `json:"content" gorm:"type:text;not null"`
}
