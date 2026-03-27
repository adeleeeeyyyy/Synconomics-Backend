package dto

type CreateBusinessRequest struct {
	Name        string  `form:"name" json:"name" validate:"required"`
	Description string  `form:"description" json:"description"`
	Category    string  `form:"category" json:"category" validate:"required"`
	Address     string  `form:"address" json:"address"`
	Latitude    float64 `form:"latitude" json:"latitude"`
	Longitude   float64 `form:"longitude" json:"longitude"`
	Phone       string  `form:"phone" json:"phone"`
	Whatsapp    string  `form:"whatsapp" json:"whatsapp"`
	Instagram   string  `form:"instagram" json:"instagram"`
	Tiktok      string  `form:"tiktok" json:"tiktok"`
	Website     string  `form:"website" json:"website"`
}

type UpdateBusinessRequest struct {
	Name        string  `form:"name" json:"name"`
	Description string  `form:"description" json:"description"`
	Category    string  `form:"category" json:"category"`
	Address     string  `form:"address" json:"address"`
	Latitude    float64 `form:"latitude" json:"latitude"`
	Longitude   float64 `form:"longitude" json:"longitude"`
	Phone       string  `form:"phone" json:"phone"`
	Whatsapp    string  `form:"whatsapp" json:"whatsapp"`
	Instagram   string  `form:"instagram" json:"instagram"`
	Tiktok      string  `form:"tiktok" json:"tiktok"`
	Website     string  `form:"website" json:"website"`
}
