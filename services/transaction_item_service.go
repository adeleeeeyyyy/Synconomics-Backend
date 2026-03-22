package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
)

type transactionItemService struct {
	repo repositories.TransactionItemRepository
}

func NewTransactionItemService(repo repositories.TransactionItemRepository) TransactionItemService {
	return &transactionItemService{repo}
}

func (s *transactionItemService) CreateTransactionItem(item *models.TransactionItem) error {
	return s.repo.Create(item)
}

func (s *transactionItemService) GetAllTransactionItems() ([]models.TransactionItem, error) {
	return s.repo.FindAll()
}

func (s *transactionItemService) GetTransactionItemById(id uint) (*models.TransactionItem, error) {
	return s.repo.FindByID(id)
}

func (s *transactionItemService) GetTransactionItemsByTransactionId(transactionID uint) ([]models.TransactionItem, error) {
	return s.repo.FindByTransactionID(transactionID)
}

func (s *transactionItemService) UpdateTransactionItem(item *models.TransactionItem) error {
	return s.repo.Update(item)
}

func (s *transactionItemService) DeleteTransactionItem(id uint) error {
	return s.repo.Delete(id)
}
