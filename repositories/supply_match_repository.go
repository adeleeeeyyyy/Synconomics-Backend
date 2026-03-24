package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type supplyMatchRepository struct {
	db *gorm.DB
}

func NewSupplyMatchRepository(db *gorm.DB) SupplyMatchRepository {
	return &supplyMatchRepository{db}
}

func (r *supplyMatchRepository) Create(match *models.SupplyMatch) error {
	return r.db.Create(match).Error
}

func (r *supplyMatchRepository) FindAll() ([]models.SupplyMatch, error) {
	var matches []models.SupplyMatch
	err := r.db.Preload("SupplyRequest").Preload("SupplyOffer").Find(&matches).Error
	return matches, err
}

func (r *supplyMatchRepository) FindByID(id uint) (*models.SupplyMatch, error) {
	var match models.SupplyMatch
	err := r.db.Preload("SupplyRequest").Preload("SupplyOffer").First(&match, id).Error
	return &match, err
}

func (r *supplyMatchRepository) FindByRequestID(requestID uint) ([]models.SupplyMatch, error) {
	var matches []models.SupplyMatch
	err := r.db.Preload("SupplyRequest").Preload("SupplyOffer").Where("supply_request_id = ?", requestID).Find(&matches).Error
	return matches, err
}

func (r *supplyMatchRepository) FindByOfferID(offerID uint) ([]models.SupplyMatch, error) {
	var matches []models.SupplyMatch
	err := r.db.Preload("SupplyRequest").Preload("SupplyOffer").Where("supply_offer_id = ?", offerID).Find(&matches).Error
	return matches, err
}

func (r *supplyMatchRepository) Update(match *models.SupplyMatch) error {
	return r.db.Save(match).Error
}

func (r *supplyMatchRepository) Delete(id uint) error {
	return r.db.Delete(&models.SupplyMatch{}, id).Error
}
