package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Business").Find(&products).Error

	return products, err
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Business").First(&product, id).Error
	return &product, err
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *productRepository) FindByBusinessID(businessID uint) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("business_id = ?", businessID).Find(&products).Error
	return products, err
}