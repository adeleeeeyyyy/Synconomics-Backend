package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
)

type productSearchLogService struct {
	repo repositories.ProductSearchLogRepository
}

func NewProductSearchLogService(repo repositories.ProductSearchLogRepository) ProductSearchLogService {
	return &productSearchLogService{repo}
}

func (s *productSearchLogService) CreateLog(log *models.ProductSearchLog) error {
	return s.repo.Create(log)
}

func (s *productSearchLogService) GetAllLogs() ([]models.ProductSearchLog, error) {
	return s.repo.FindAll()
}

func (s *productSearchLogService) GetLogById(id uint) (*models.ProductSearchLog, error) {
	return s.repo.FindByID(id)
}

func (s *productSearchLogService) GetLogsByUserId(userID uint) ([]models.ProductSearchLog, error) {
	return s.repo.FindByUserID(userID)
}
