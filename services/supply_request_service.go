package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
)

type supplyRequestService struct {
	repo repositories.SupplyRequestRepository
}

func NewSupplyRequestService(repo repositories.SupplyRequestRepository) SupplyRequestService {
	return &supplyRequestService{repo}
}

func (s *supplyRequestService) CreateSupplyRequest(request *models.SupplyRequest) error {
	return s.repo.Create(request)
}

func (s *supplyRequestService) GetAllSupplyRequests() ([]models.SupplyRequest, error) {
	return s.repo.FindAll()
}

func (s *supplyRequestService) GetSupplyRequestById(id uint) (*models.SupplyRequest, error) {
	return s.repo.FindByID(id)
}

func (s *supplyRequestService) GetSupplyRequestsByBusinessId(businessID uint) ([]models.SupplyRequest, error) {
	return s.repo.FindByBusinessID(businessID)
}

func (s *supplyRequestService) UpdateSupplyRequest(request *models.SupplyRequest) error {
	return s.repo.Update(request)
}

func (s *supplyRequestService) DeleteSupplyRequest(id uint) error {
	return s.repo.Delete(id)
}
