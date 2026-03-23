package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type supplyOfferRepository struct {
	db *gorm.DB
}

func NewSupplyOfferRepository(db *gorm.DB) SupplyOfferRepository {
	return &supplyOfferRepository{db}
}

func (r *supplyOfferRepository) Create(offer *models.SupplyOffer) error {
	return r.db.Create(offer).Error
}

func (r *supplyOfferRepository) FindAll() ([]models.SupplyOffer, error) {
	var offers []models.SupplyOffer
	err := r.db.Preload("Business").Preload("Product").Find(&offers).Error
	return offers, err
}

func (r *supplyOfferRepository) FindByID(id uint) (*models.SupplyOffer, error) {
	var offer models.SupplyOffer
	err := r.db.Preload("Business").Preload("Product").First(&offer, id).Error
	return &offer, err
}

func (r *supplyOfferRepository) FindByBusinessID(businessID uint) ([]models.SupplyOffer, error) {
	var offers []models.SupplyOffer
	err := r.db.Preload("Business").Preload("Product").Where("business_id = ?", businessID).Find(&offers).Error
	return offers, err
}

func (r *supplyOfferRepository) Update(offer *models.SupplyOffer) error {
	return r.db.Save(offer).Error
}

func (r *supplyOfferRepository) Delete(id uint) error {
	return r.db.Delete(&models.SupplyOffer{}, id).Error
}
