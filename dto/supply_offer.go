package dto

type CreateSupplyOfferReq struct {
	BusinessID  uint   `json:"business_id" validate:"required"`
	ProductID   uint   `json:"product_id"`
	ProductName string `json:"product_name" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required,min=1"`
}

type UpdateSupplyOfferReq struct {
	ProductID   uint   `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}
