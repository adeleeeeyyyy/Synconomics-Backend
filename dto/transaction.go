package dto

type CreateTransactionRequest struct {
	BusinessID    uint    `json:"business_id" validate:"required"`
	TotalAmount   float64 `json:"total_amount" validate:"required"`
	PaymentMethod string  `json:"payment_method" validate:"required"`
	Status        string  `json:"status" validate:"required,oneof=pending completed cancelled"`
}

type UpdateTransactionRequest struct {
	TotalAmount   float64 `json:"total_amount"`
	PaymentMethod string  `json:"payment_method"`
	Status        string  `json:"status" validate:"omitempty,oneof=pending completed cancelled"`
}
