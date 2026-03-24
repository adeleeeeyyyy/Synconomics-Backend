package dto

type CreateSupplyRequestReq struct {
	BusinessID  uint   `json:"business_id" validate:"required"`
	ProductName string `json:"product_name" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required,min=1"`
	Status      string `json:"status" validate:"required,oneof=open matched closed"`
}

type UpdateSupplyRequestReq struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Status      string `json:"status" validate:"omitempty,oneof=open matched closed"`
}
