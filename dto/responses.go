package dto

import "time"

// Auth & User
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	GoogleID  string    `json:"google_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// Business Management
type BusinessResponse struct {
	ID          uint    `json:"id"`
	UserID      uint    `json:"user_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Address     string  `json:"address"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	Phone       string  `form:"phone" json:"phone"`
	Whatsapp    string  `form:"whatsapp" json:"whatsapp"`
	Instagram   string  `form:"instagram" json:"instagram"`
	Tiktok      string  `form:"tiktok" json:"tiktok"`
	Website     string  `form:"website" json:"website"`
}

type ProductResponse struct {
	ID          uint      `json:"id"`
	BusinessID  uint      `json:"business_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	ImageURL    string    `json:"image_url"`
	MinStock    int       `json:"min_stock"`
	CreatedAt   time.Time `json:"created_at"`
}

type TransactionResponse struct {
	ID              uint      `json:"id"`
	BusinessID      uint      `json:"business_id"`
	Type            string    `json:"type"`
	TotalAmount     float64   `json:"total_amount"`
	TransactionDate time.Time `json:"transaction_date"`
}

type ExpenseResponse struct {
	ID          uint      `json:"id"`
	BusinessID  uint      `json:"business_id"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	ExpenseDate time.Time `json:"expense_date"`
}

// AI Features
type AISessionResponse struct {
	ID         uint      `json:"id"`
	BusinessID uint      `json:"business_id"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
}

type AIMessageResponse struct {
	ID        uint      `json:"id"`
	SessionID uint      `json:"session_id"`
	Role      string    `json:"role"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type AIResultResponse struct {
	ID        uint      `json:"id"`
	SessionID uint      `json:"session_id"`
	Result    string    `json:"result"`
	CreatedAt time.Time `json:"created_at"`
}

// Supply Chain
type SupplyRequestResponse struct {
	ID          uint      `json:"id"`
	BusinessID  uint      `json:"business_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type SupplyOfferResponse struct {
	ID          uint      `json:"id"`
	BusinessID  uint      `json:"business_id"`
	ProductName string    `json:"product_name"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type SupplyMatchResponse struct {
	ID              uint      `json:"id"`
	SupplyRequestID uint      `json:"supply_request_id"`
	SupplyOfferID   uint      `json:"supply_offer_id"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}

// Community
type ThreadResponse struct {
	ID        uint          `json:"id"`
	UserID    uint          `json:"user_id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
	User      *UserResponse `json:"user,omitempty"`
}

type ReplyResponse struct {
	ID        uint          `json:"id"`
	ThreadID  uint          `json:"thread_id"`
	UserID    uint          `json:"user_id"`
	Content   string        `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
	User      *UserResponse `json:"user,omitempty"`
}

// Analytics/Logs
type ProductSearchLogResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Keyword   string    `json:"keyword"`
	CreatedAt time.Time `json:"created_at"`
}
