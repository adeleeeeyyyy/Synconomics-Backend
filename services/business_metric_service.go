package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
)

type businessMetricService struct {
	repo repositories.BusinessMetricRepository
}

func NewBusinessMetricService(repo repositories.BusinessMetricRepository) BusinessMetricService {
	return &businessMetricService{repo}
}

func (s *businessMetricService) CreateMetric(metric *models.BusinessMetric) error {
	return s.repo.Create(metric)
}

func (s *businessMetricService) GetAllMetrics() ([]models.BusinessMetric, error) {
	return s.repo.FindAll()
}

func (s *businessMetricService) GetMetricById(id uint) (*models.BusinessMetric, error) {
	return s.repo.FindByID(id)
}

func (s *businessMetricService) GetMetricsByBusinessId(businessID uint) ([]models.BusinessMetric, error) {
	return s.repo.FindByBusinessID(businessID)
}

func (s *businessMetricService) UpdateMetric(metric *models.BusinessMetric) error {
	return s.repo.Update(metric)
}

func (s *businessMetricService) DeleteMetric(id uint) error {
	return s.repo.Delete(id)
}

func (s *businessMetricService) GetLatestMetricByBusinessId(businessID uint) (*models.BusinessMetric, error) {
	return s.repo.GetLatestByBusinessID(businessID)
}
