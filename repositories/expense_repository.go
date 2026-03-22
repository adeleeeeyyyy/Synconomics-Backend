package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type expenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepository{db}
}

func (r *expenseRepository) Create(expense *models.Expense) error {
	return r.db.Create(expense).Error
}

func (r *expenseRepository) FindAll() ([]models.Expense, error) {
	var expenses []models.Expense
	err := r.db.Preload("Business").Find(&expenses).Error
	return expenses, err
}

func (r *expenseRepository) FindByID(id uint) (*models.Expense, error) {
	var expense models.Expense
	err := r.db.Preload("Business").First(&expense, id).Error
	return &expense, err
}

func (r *expenseRepository) FindByBusinessID(businessID uint) ([]models.Expense, error) {
	var expenses []models.Expense
	err := r.db.Preload("Business").Where("business_id = ?", businessID).Find(&expenses).Error
	return expenses, err
}

func (r *expenseRepository) Update(expense *models.Expense) error {
	return r.db.Save(expense).Error
}

func (r *expenseRepository) Delete(id uint) error {
	return r.db.Delete(&models.Expense{}, id).Error
}
