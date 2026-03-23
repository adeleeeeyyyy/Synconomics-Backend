package dto

type CreateSupplyMatchReq struct {
	SupplyRequestID uint   `json:"supply_request_id" validate:"required"`
	SupplyOfferID   uint   `json:"supply_offer_id" validate:"required"`
	Status          string `json:"status" validate:"omitempty,oneof=pending accepted rejected"`
}

type UpdateSupplyMatchStatusReq struct {
	Status string `json:"status" validate:"required,oneof=pending accepted rejected"`
}
