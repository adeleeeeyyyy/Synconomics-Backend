package models

import "gorm.io/gorm"

// MatchStatus defines the enum for the status field
type MatchStatus string

const (
	StatusMatchPending  MatchStatus = "pending"
	StatusAccepted MatchStatus = "accepted"
	StatusRejected MatchStatus = "rejected"
)

type SupplyMatch struct {
	gorm.Model
	SupplyRequestID uint          `json:"request_id"`
	SupplyRequest   SupplyRequest `json:"supply_request" gorm:"foreignKey:SupplyRequestID"`

	SupplyOfferID uint        `json:"offer_id"`
	SupplyOffer   SupplyOffer `json:"supply_offer" gorm:"foreignKey:SupplyOfferID"`

	Status MatchStatus `json:"status" gorm:"type:enum('pending', 'accepted', 'rejected');default:'pending'"`
}

