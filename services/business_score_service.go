package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
)

type businessScoreService struct {
	repo repositories.BusinessScoreRepository
}

func NewBusinessScoreService(repo repositories.BusinessScoreRepository) BusinessScoreService {
	return &businessScoreService{repo}
}

func (s *businessScoreService) CreateScore(score *models.BusinessScore) error {
	return s.repo.Create(score)
}

func (s *businessScoreService) GetAllScores() ([]models.BusinessScore, error) {
	return s.repo.FindAll()
}

func (s *businessScoreService) GetScoreById(id uint) (*models.BusinessScore, error) {
	return s.repo.FindByID(id)
}

func (s *businessScoreService) GetScoresByBusinessId(businessID uint) ([]models.BusinessScore, error) {
	return s.repo.FindByBusinessID(businessID)
}

func (s *businessScoreService) UpdateScore(score *models.BusinessScore) error {
	return s.repo.Update(score)
}

func (s *businessScoreService) DeleteScore(id uint) error {
	return s.repo.Delete(id)
}

func (s *businessScoreService) GetLatestScoreByBusinessId(businessID uint) (*models.BusinessScore, error) {
	return s.repo.GetLatestByBusinessID(businessID)
}
