package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
)

type marketTrendService struct {
	repo repositories.MarketTrendRepository
}

func NewMarketTrendService(repo repositories.MarketTrendRepository) MarketTrendService {
	return &marketTrendService{repo}
}

func (s *marketTrendService) CreateTrend(trend *models.MarketTrend) error {
	return s.repo.Create(trend)
}

func (s *marketTrendService) GetAllTrends() ([]models.MarketTrend, error) {
	return s.repo.FindAll()
}

func (s *marketTrendService) GetTrendById(id uint) (*models.MarketTrend, error) {
	return s.repo.FindByID(id)
}

func (s *marketTrendService) UpdateTrend(trend *models.MarketTrend) error {
	return s.repo.Update(trend)
}

func (s *marketTrendService) DeleteTrend(id uint) error {
	return s.repo.Delete(id)
}

func (s *marketTrendService) GetTopTenTrends() ([]models.MarketTrend, error) {
	return s.repo.FindTopTen()
}
