package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
)

type threadService struct {
	repo repositories.ThreadRepository
}

func NewThreadService(repo repositories.ThreadRepository) ThreadService {
	return &threadService{repo}
}

func (s *threadService) CreateThread(thread *models.Thread) error {
	return s.repo.Create(thread)
}

func (s *threadService) GetAllThreads() ([]models.Thread, error) {
	return s.repo.FindAll()
}

func (s *threadService) GetThreadById(id uint) (*models.Thread, error) {
	return s.repo.FindByID(id)
}

func (s *threadService) UpdateThread(thread *models.Thread) error {
	return s.repo.Update(thread)
}

func (s *threadService) DeleteThread(id uint) error {
	return s.repo.Delete(id)
}
