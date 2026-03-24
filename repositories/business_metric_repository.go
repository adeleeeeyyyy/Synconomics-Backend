package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type businessMetricRepository struct {
	db *gorm.DB
}

func NewBusinessMetricRepository(db *gorm.DB) BusinessMetricRepository {
	return &businessMetricRepository{db}
}

func (r *businessMetricRepository) Create(metric *models.BusinessMetric) error {
	return r.db.Create(metric).Error
}

func (r *businessMetricRepository) FindAll() ([]models.BusinessMetric, error) {
	var metrics []models.BusinessMetric
	err := r.db.Preload("Business").Find(&metrics).Error
	return metrics, err
}

func (r *businessMetricRepository) FindByID(id uint) (*models.BusinessMetric, error) {
	var metric models.BusinessMetric
	err := r.db.Preload("Business").First(&metric, id).Error
	return &metric, err
}

func (r *businessMetricRepository) FindByBusinessID(businessID uint) ([]models.BusinessMetric, error) {
	var metrics []models.BusinessMetric
	err := r.db.Preload("Business").Where("business_id = ?", businessID).Find(&metrics).Error
	return metrics, err
}

func (r *businessMetricRepository) Update(metric *models.BusinessMetric) error {
	return r.db.Save(metric).Error
}

func (r *businessMetricRepository) Delete(id uint) error {
	return r.db.Delete(&models.BusinessMetric{}, id).Error
}

func (r *businessMetricRepository) GetLatestByBusinessID(businessID uint) (*models.BusinessMetric, error) {
	var metric models.BusinessMetric
	err := r.db.Preload("Business").Where("business_id = ?", businessID).Order("calculated_at desc").First(&metric).Error
	return &metric, err
}
