package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type supplyRequestRepository struct {
	db *gorm.DB
}

func NewSupplyRequestRepository(db *gorm.DB) SupplyRequestRepository {
	return &supplyRequestRepository{db}
}

func (r *supplyRequestRepository) Create(request *models.SupplyRequest) error {
	return r.db.Create(request).Error
}

func (r *supplyRequestRepository) FindAll() ([]models.SupplyRequest, error) {
	var requests []models.SupplyRequest
	err := r.db.Preload("Business").Find(&requests).Error
	return requests, err
}

func (r *supplyRequestRepository) FindByID(id uint) (*models.SupplyRequest, error) {
	var request models.SupplyRequest
	err := r.db.Preload("Business").First(&request, id).Error
	return &request, err
}

func (r *supplyRequestRepository) FindByBusinessID(businessID uint) ([]models.SupplyRequest, error) {
	var requests []models.SupplyRequest
	err := r.db.Preload("Business").Where("business_id = ?", businessID).Find(&requests).Error
	return requests, err
}

func (r *supplyRequestRepository) Update(request *models.SupplyRequest) error {
	return r.db.Save(request).Error
}

func (r *supplyRequestRepository) Delete(id uint) error {
	return r.db.Delete(&models.SupplyRequest{}, id).Error
}
