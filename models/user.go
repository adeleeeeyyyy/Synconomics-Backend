package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	AISessions        []AISession        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Replies           []Reply            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Threads           []Thread           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductSearchLogs []ProductSearchLog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Name     string  `json:"name"                    gorm:"type:varchar(100);not null"`
	Email    string  `json:"email"                   gorm:"type:varchar(100);uniqueIndex;not null"`
	Password *string `json:"-"                       gorm:"type:varchar(255)"`
	Provider string  `json:"provider"                gorm:"type:varchar(50);not null;default:'manual'"`
	GoogleID *string `json:"google_id,omitempty"     gorm:"type:varchar(255);uniqueIndex"`
	Avatar   *string `json:"avatar,omitempty"        gorm:"type:varchar(255)"`
}
