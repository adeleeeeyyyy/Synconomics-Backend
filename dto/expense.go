package dto

type CreateExpenseRequest struct {
	BusinessID uint    `json:"business_id" validate:"required"`
	Title      string  `json:"title" validate:"required"`
	Amount     float64 `json:"amount" validate:"required"`
	Category   string  `json:"category" validate:"required"`
}

type UpdateExpenseRequest struct {
	Title    string  `json:"title"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
}
