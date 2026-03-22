package dto

type CreateProductRequest struct {
	BusinessID  uint   `json:"business_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Price       int    `json:"price" validate:"required, min=1000"`
	Stock       int    `json:"stock" validate:"required"`
	MinStock    int    `json:"min_stock"`
}
