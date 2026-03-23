package dto

type CreateThreadReq struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type UpdateThreadReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
