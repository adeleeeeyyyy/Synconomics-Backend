package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
)

type businessService struct {
	businessRepo repositories.BusinessRepository
}

func NewBusinessService(repo repositories.BusinessRepository) BusinessService {
	return &businessService{
		businessRepo: repo,
	}
}

func (s *businessService) CreateBusiness(business *models.Business) error {
	return s.businessRepo.Create(business)
}

func (s *businessService) GetAllBusinesses() ([]models.Business, error) {
	return s.businessRepo.FindAll()
}

func (s *businessService) GetBusinessById(id uint) (*models.Business, error) {
	return s.businessRepo.FindByID(id)
}

func (s *businessService) UpdateBusiness(business *models.Business) error {
	return s.businessRepo.Update(business)
}

func (s *businessService) DeleteBusiness(id uint) error {
	return s.businessRepo.Delete(id)
}