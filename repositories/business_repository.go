package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type businessRepository struct {
	db *gorm.DB
}

func NewBusinessRepository(db *gorm.DB) BusinessRepository {
	return &businessRepository{db}
}

func (r *businessRepository) Create(business *models.Business) error {
	return r.db.Create(business).Error
}

func (r *businessRepository) FindAll() ([]models.Business, error) {
	var businesses []models.Business
	err := r.db.Preload("User").Find(&businesses).Error
	return businesses, err
}

func (r *businessRepository) FindByID(id uint) (*models.Business, error) {
	var business models.Business
	err := r.db.Preload("User").First(&business, id).Error
	if err != nil {
		return nil, err
	}
	return &business, nil
}

func (r *businessRepository) Update(business *models.Business) error {
	return r.db.Save(business).Error
}

func (r *businessRepository) Delete(id uint) error {
	return r.db.Delete(&models.Business{}, id).Error
}

func (r *businessRepository) FindByUserID(userID uint) ([]models.Business, error) {
	var businesses []models.Business
	err := r.db.Where("user_id = ?", userID).Find(&businesses).Error
	return businesses, err
}