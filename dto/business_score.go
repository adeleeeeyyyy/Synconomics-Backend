package dto

import "time"

type CreateBusinessScoreRequest struct {
	BusinessID uint   `json:"business_id" validate:"required"`
	Score      int    `json:"score" validate:"required"`
	Insight    string `json:"insight"`
}

type UpdateBusinessScoreRequest struct {
	Score   int    `json:"score"`
	Insight string `json:"insight"`
}

type BusinessScoreResponse struct {
	ID         uint      `json:"id"`
	BusinessID uint      `json:"business_id"`
	Score      int       `json:"score"`
	Insight    string    `json:"insight"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
