package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `json:"name"                    gorm:"type:varchar(100);not null"`
	Email    string  `json:"email"                   gorm:"type:varchar(100);uniqueIndex;not null"`
	Password *string `json:"-"                       gorm:"type:varchar(255)"`
	Provider string  `json:"provider"                gorm:"type:varchar(50);not null;default:'manual'"`
	GoogleID *string `json:"google_id,omitempty"     gorm:"type:varchar(255);uniqueIndex"`
	Avatar   *string `json:"avatar,omitempty"        gorm:"type:varchar(255)"`
}
