package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
	"fmt"
)

type transactionItemService struct {
	repo        repositories.TransactionItemRepository
	productRepo repositories.ProductRepository
}

func NewTransactionItemService(repo repositories.TransactionItemRepository, productRepo repositories.ProductRepository) TransactionItemService {
	return &transactionItemService{repo, productRepo}
}

func (s *transactionItemService) CreateTransactionItem(item *models.TransactionItem) error {
	// 1. Get Product
	product, err := s.productRepo.FindByID(item.ProductID)
	if err != nil {
		return err
	}

	// 2. Check Stock
	if product.Stock < item.Quantity {
		return fmt.Errorf("insufficient stock for product %s (available: %d, requested: %d)", product.Name, product.Stock, item.Quantity)
	}

	// 3. Decrease Stock
	product.Stock -= item.Quantity
	if err := s.productRepo.Update(product); err != nil {
		return err
	}

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
	// 1. Get Existing Item
	existingItem, err := s.repo.FindByID(item.ID)
	if err != nil {
		return err
	}

	// 2. Get Product
	product, err := s.productRepo.FindByID(item.ProductID)
	if err != nil {
		return err
	}

	// 3. Adjust Stock
	stockAdjustment := existingItem.Quantity - item.Quantity
	if product.Stock+stockAdjustment < 0 {
		return fmt.Errorf("insufficient stock for adjustment (available: %d, requested adjustment: %d)", product.Stock, -stockAdjustment)
	}

	product.Stock += stockAdjustment
	if err := s.productRepo.Update(product); err != nil {
		return err
	}

	return s.repo.Update(item)
}

func (s *transactionItemService) DeleteTransactionItem(id uint) error {
	// 1. Get Existing Item
	existingItem, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// 2. Get Product
	product, err := s.productRepo.FindByID(existingItem.ProductID)
	if err == nil { // Restore stock if product still exists
		product.Stock += existingItem.Quantity
		s.productRepo.Update(product)
	}

	return s.repo.Delete(id)
}
