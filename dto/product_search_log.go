package dto

type CreateProductSearchLogReq struct {
	Keyword string `json:"keyword" validate:"required"`
}
