package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
)

type expenseService struct {
	repo repositories.ExpenseRepository
}

func NewExpenseService(repo repositories.ExpenseRepository) ExpenseService {
	return &expenseService{repo}
}

func (s *expenseService) CreateExpense(expense *models.Expense) error {
	return s.repo.Create(expense)
}

func (s *expenseService) GetAllExpenses() ([]models.Expense, error) {
	return s.repo.FindAll()
}

func (s *expenseService) GetExpenseById(id uint) (*models.Expense, error) {
	return s.repo.FindByID(id)
}

func (s *expenseService) GetExpensesByBusinessId(businessID uint) ([]models.Expense, error) {
	return s.repo.FindByBusinessID(businessID)
}

func (s *expenseService) UpdateExpense(expense *models.Expense) error {
	return s.repo.Update(expense)
}

func (s *expenseService) DeleteExpense(id uint) error {
	return s.repo.Delete(id)
}
