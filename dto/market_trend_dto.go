package dto

type CreateMarketTrendRequest struct {
	ProductName string  `json:"product_name" validate:"required"`
	SearchCount int     `json:"search_count"`
	DemandScore float64 `json:"demand_score"`
	Location    string  `json:"location"`
}

type UpdateMarketTrendRequest struct {
	ProductName string  `json:"product_name"`
	SearchCount int     `json:"search_count"`
	DemandScore float64 `json:"demand_score"`
	Location    string  `json:"location"`
}
