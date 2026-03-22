package dto

type CreateTransactionItemRequest struct {
	TransactionID uint    `json:"transaction_id" validate:"required"`
	ProductID     uint    `json:"product_id" validate:"required"`
	Quantity      int     `json:"quantity" validate:"required,min=1"`
	Price         float64 `json:"price" validate:"required"`
}

type UpdateTransactionItemRequest struct {
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
