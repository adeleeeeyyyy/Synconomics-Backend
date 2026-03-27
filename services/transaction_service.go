package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
	"time"
)

type transactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(repo repositories.TransactionRepository) TransactionService {
	return &transactionService{repo}
}

func (s *transactionService) CreateTransaction(transaction *models.Transaction) error {
	return s.repo.Create(transaction)
}

func (s *transactionService) GetAllTransactions() ([]models.Transaction, error) {
	return s.repo.FindAll()
}

func (s *transactionService) GetTransactionById(id uint) (*models.Transaction, error) {
	return s.repo.FindByID(id)
}

func (s *transactionService) UpdateTransaction(transaction *models.Transaction) error {
	return s.repo.Update(transaction)
}

func (s *transactionService) DeleteTransaction(id uint) error {
	return s.repo.Delete(id)
}

func (s *transactionService) GetTransactionsByBusinessId(businessID uint) ([]models.Transaction, error) {
	return s.repo.FindByBusinessID(businessID)
}

func (s *transactionService) GetTransactionsByDateRange(businessID uint, startDate, endDate time.Time) ([]models.Transaction, error) {
	return s.repo.FindByBusinessIDAndDateRange(businessID, startDate, endDate)
}
