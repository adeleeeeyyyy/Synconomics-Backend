package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type businessScoreRepository struct {
	db *gorm.DB
}

func NewBusinessScoreRepository(db *gorm.DB) BusinessScoreRepository {
	return &businessScoreRepository{db}
}

func (r *businessScoreRepository) Create(score *models.BusinessScore) error {
	return r.db.Create(score).Error
}

func (r *businessScoreRepository) FindAll() ([]models.BusinessScore, error) {
	var scores []models.BusinessScore
	err := r.db.Find(&scores).Error
	return scores, err
}

func (r *businessScoreRepository) FindByID(id uint) (*models.BusinessScore, error) {
	var score models.BusinessScore
	err := r.db.First(&score, id).Error
	return &score, err
}

func (r *businessScoreRepository) FindByBusinessID(businessID uint) ([]models.BusinessScore, error) {
	var scores []models.BusinessScore
	err := r.db.Where("business_id = ?", businessID).Find(&scores).Error
	return scores, err
}

func (r *businessScoreRepository) Update(score *models.BusinessScore) error {
	return r.db.Save(score).Error
}

func (r *businessScoreRepository) Delete(id uint) error {
	return r.db.Delete(&models.BusinessScore{}, id).Error
}

func (r *businessScoreRepository) GetLatestByBusinessID(businessID uint) (*models.BusinessScore, error) {
	var score models.BusinessScore
	err := r.db.Where("business_id = ?", businessID).Order("created_at desc").First(&score).Error
	return &score, err
}
