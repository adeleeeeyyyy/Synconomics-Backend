package dto

type CreateReplyReq struct {
	ThreadID uint   `json:"thread_id" validate:"required"`
	Content  string `json:"content" validate:"required"`
}

type UpdateReplyReq struct {
	Content string `json:"content" validate:"required"`
}
