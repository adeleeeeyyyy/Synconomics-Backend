package models

import "gorm.io/gorm"

type ProductSearchLog struct {
	gorm.Model
	UserID uint `json:"user_id"`
	User   User `json:"user"`

	Keyword string `json:"keyword" gorm:"type:varchar(255)"`
}
