package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
)

type supplyOfferService struct {
	repo repositories.SupplyOfferRepository
}

func NewSupplyOfferService(repo repositories.SupplyOfferRepository) SupplyOfferService {
	return &supplyOfferService{repo}
}

func (s *supplyOfferService) CreateSupplyOffer(offer *models.SupplyOffer) error {
	return s.repo.Create(offer)
}

func (s *supplyOfferService) GetAllSupplyOffers() ([]models.SupplyOffer, error) {
	return s.repo.FindAll()
}

func (s *supplyOfferService) GetSupplyOfferById(id uint) (*models.SupplyOffer, error) {
	return s.repo.FindByID(id)
}

func (s *supplyOfferService) GetSupplyOffersByBusinessId(businessID uint) ([]models.SupplyOffer, error) {
	return s.repo.FindByBusinessID(businessID)
}

func (s *supplyOfferService) UpdateSupplyOffer(offer *models.SupplyOffer) error {
	return s.repo.Update(offer)
}

func (s *supplyOfferService) DeleteSupplyOffer(id uint) error {
	return s.repo.Delete(id)
}
