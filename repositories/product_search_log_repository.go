package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type productSearchLogRepository struct {
	db *gorm.DB
}

func NewProductSearchLogRepository(db *gorm.DB) ProductSearchLogRepository {
	return &productSearchLogRepository{db}
}

func (r *productSearchLogRepository) Create(log *models.ProductSearchLog) error {
	return r.db.Create(log).Error
}

func (r *productSearchLogRepository) FindAll() ([]models.ProductSearchLog, error) {
	var logs []models.ProductSearchLog
	err := r.db.Preload("User").Find(&logs).Error
	return logs, err
}

func (r *productSearchLogRepository) FindByID(id uint) (*models.ProductSearchLog, error) {
	var log models.ProductSearchLog
	err := r.db.Preload("User").First(&log, id).Error
	return &log, err
}

func (r *productSearchLogRepository) FindByUserID(userID uint) ([]models.ProductSearchLog, error) {
	var logs []models.ProductSearchLog
	err := r.db.Preload("User").Where("user_id = ?", userID).Find(&logs).Error
	return logs, err
}
