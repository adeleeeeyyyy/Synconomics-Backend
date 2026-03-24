package dto

// CreateSessionRequest mendefinisikan payload untuk pembuatan AI Session
type CreateSessionRequest struct {
	BusinessID uint   `json:"business_id" validate:"required"`
	Type       string `json:"type" validate:"required,oneof=idea_generation validation strategy"`
}

// ChatRequest mendefinisikan payload untuk mengirim pesan ke Gemini
type ChatRequest struct {
	Message string `json:"message" validate:"required"`
}

// ChatWithRoleRequest mendefinisikan payload untuk chat dengan role tertentu tanpa session_id di URL
type ChatWithRoleRequest struct {
	BusinessID uint   `json:"business_id" validate:"required"`
	Message    string `json:"message" validate:"required"`
}
