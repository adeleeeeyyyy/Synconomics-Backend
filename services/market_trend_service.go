package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
)

type marketTrendService struct {
	repo       repositories.MarketTrendRepository
	aiService  AIService
	logRepo    repositories.ProductSearchLogRepository
}

func NewMarketTrendService(
	repo repositories.MarketTrendRepository,
	aiService AIService,
	logRepo repositories.ProductSearchLogRepository,
) MarketTrendService {
	return &marketTrendService{repo, aiService, logRepo}
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

func (s *marketTrendService) RefreshTrendsFromLogs() error {
	// 1. Get recent unique keywords from logs
	keywords, err := s.logRepo.GetRecentUniqueKeywords(100) // process top 100 recent keywords
	if err != nil {
		return err
	}

	if len(keywords) == 0 {
		return nil
	}

	// 2. Analyze with AI
	trends, err := s.aiService.AnalyzeMarketTrends(keywords)
	if err != nil {
		return err
	}

	// 3. Replace old trends with new ones
	return s.repo.ReplaceAll(trends)
}
