package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
)

type supplyMatchService struct {
	repo repositories.SupplyMatchRepository
}

func NewSupplyMatchService(repo repositories.SupplyMatchRepository) SupplyMatchService {
	return &supplyMatchService{repo}
}

func (s *supplyMatchService) CreateSupplyMatch(match *models.SupplyMatch) error {
	return s.repo.Create(match)
}

func (s *supplyMatchService) GetAllSupplyMatches() ([]models.SupplyMatch, error) {
	return s.repo.FindAll()
}

func (s *supplyMatchService) GetSupplyMatchById(id uint) (*models.SupplyMatch, error) {
	return s.repo.FindByID(id)
}

func (s *supplyMatchService) GetSupplyMatchesByRequestId(requestID uint) ([]models.SupplyMatch, error) {
	return s.repo.FindByRequestID(requestID)
}

func (s *supplyMatchService) GetSupplyMatchesByOfferId(offerID uint) ([]models.SupplyMatch, error) {
	return s.repo.FindByOfferID(offerID)
}

func (s *supplyMatchService) UpdateSupplyMatch(match *models.SupplyMatch) error {
	return s.repo.Update(match)
}

func (s *supplyMatchService) DeleteSupplyMatch(id uint) error {
	return s.repo.Delete(id)
}
