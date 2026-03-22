package models

import "gorm.io/gorm"

type AISessionType string

const (
	IdeaGeneration AISessionType = "idea_generation"
	Validation     AISessionType = "validation"
	Strategy       AISessionType = "strategy"
)

type AISession struct {
	gorm.Model
	UserID	 uint
	BusinessID	uint
	Business	Business
	User	 User
	AIMessages []AIMessage `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AIResults  []AIResult  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Type     AISessionType `gorm:"type:enum('idea_generation','validation','strategy');" json:"type"`
}
