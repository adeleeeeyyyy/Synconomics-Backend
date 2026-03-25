package dto

type CreateProductRequest struct {
	BusinessID  uint   `json:"business_id" form:"business_id" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price" validate:"required,min=1000"`
	Stock       int    `json:"stock" form:"stock" validate:"required"`
	MinStock    int    `json:"min_stock" form:"min_stock"`
}
type UpdateProductRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	Stock       int    `json:"stock" form:"stock"`
	MinStock    int    `json:"min_stock" form:"min_stock"`
}
