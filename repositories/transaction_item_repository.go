package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type transactionItemRepository struct {
	db *gorm.DB
}

func NewTransactionItemRepository(db *gorm.DB) TransactionItemRepository {
	return &transactionItemRepository{db}
}

func (r *transactionItemRepository) Create(item *models.TransactionItem) error {
	return r.db.Create(item).Error
}

func (r *transactionItemRepository) FindAll() ([]models.TransactionItem, error) {
	var items []models.TransactionItem
	err := r.db.Preload("Product").Find(&items).Error
	return items, err
}

func (r *transactionItemRepository) FindByID(id uint) (*models.TransactionItem, error) {
	var item models.TransactionItem
	err := r.db.Preload("Product").First(&item, id).Error
	return &item, err
}

func (r *transactionItemRepository) FindByTransactionID(transactionID uint) ([]models.TransactionItem, error) {
	var items []models.TransactionItem
	err := r.db.Preload("Product").Where("transaction_id = ?", transactionID).Find(&items).Error
	return items, err
}

func (r *transactionItemRepository) Update(item *models.TransactionItem) error {
	return r.db.Save(item).Error
}

func (r *transactionItemRepository) Delete(id uint) error {
	return r.db.Delete(&models.TransactionItem{}, id).Error
}
