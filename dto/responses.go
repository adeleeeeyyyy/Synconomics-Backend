package dto

import "time"

// Auth & User
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	GoogleID  string    `json:"google_id,omitempty"`
	Avatar    *string   `json:"avatar,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserBusinessResponse struct {
	User       UserResponse       `json:"user"`
	Businesses []BusinessResponse `json:"businesses"`
}

// Business Management
type BusinessResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	LogoURL     string    `json:"logo_url"`
	Address     string    `json:"address"`
	Latitude    float32   `json:"latitude"`
	Longitude   float32   `json:"longitude"`
	Phone       string    `json:"phone"`
	Whatsapp    string    `json:"whatsapp"`
	Instagram   string    `json:"instagram"`
	Tiktok      string    `json:"tiktok"`
	Website     string    `json:"website"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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
	UpdatedAt   time.Time `json:"updated_at"`
}

type TransactionResponse struct {
	ID              uint      `json:"id"`
	BusinessID      uint      `json:"business_id"`
	TotalAmount     float64   `json:"total_amount"`
	PaymentMethod   string    `json:"payment_method"`
	Status          string    `json:"status"`
	TransactionDate time.Time `json:"transaction_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ExpenseResponse struct {
	ID          uint      `json:"id"`
	BusinessID  uint      `json:"business_id"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
	Title       string    `json:"title"`
	ExpenseDate time.Time `json:"expense_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// AI Features
type AISessionResponse struct {
	ID         uint      `json:"id"`
	BusinessID uint      `json:"business_id"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type AIMessageResponse struct {
	ID        uint      `json:"id"`
	SessionID uint      `json:"session_id"`
	Role      string    `json:"role"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AIResultResponse struct {
	ID        uint      `json:"id"`
	SessionID uint      `json:"session_id"`
	Result    string    `json:"result"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Supply Chain
type SupplyRequestResponse struct {
	ID          uint      `json:"id"`
	BusinessID  uint      `json:"business_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SupplyOfferResponse struct {
	ID          uint      `json:"id"`
	BusinessID  uint      `json:"business_id"`
	ProductID   uint      `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SupplyMatchResponse struct {
	ID              uint      `json:"id"`
	SupplyRequestID uint      `json:"supply_request_id"`
	SupplyOfferID   uint      `json:"supply_offer_id"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Community
type ThreadResponse struct {
	ID        uint          `json:"id"`
	UserID    uint          `json:"user_id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      *UserResponse `json:"user,omitempty"`
}

type ReplyResponse struct {
	ID        uint          `json:"id"`
	ThreadID  uint          `json:"thread_id"`
	UserID    uint          `json:"user_id"`
	Content   string        `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      *UserResponse `json:"user,omitempty"`
}

// Analytics/Logs
type ProductSearchLogResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Keyword   string    `json:"keyword"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MarketTrendResponse struct {
	ID          uint      `json:"id"`
	ProductName string    `json:"product_name"`
	SearchCount int       `json:"search_count"`
	DemandScore float64   `json:"demand_score"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
