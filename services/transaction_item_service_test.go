package services

import (
	"Synconomics/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionItemRepository struct {
	mock.Mock
}

func (m *MockTransactionItemRepository) Create(item *models.TransactionItem) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockTransactionItemRepository) FindAll() ([]models.TransactionItem, error) {
	args := m.Called()
	return args.Get(0).([]models.TransactionItem), args.Error(1)
}

func (m *MockTransactionItemRepository) FindByID(id uint) (*models.TransactionItem, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.TransactionItem), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTransactionItemRepository) FindByTransactionID(transactionID uint) ([]models.TransactionItem, error) {
	args := m.Called(transactionID)
	return args.Get(0).([]models.TransactionItem), args.Error(1)
}

func (m *MockTransactionItemRepository) Update(item *models.TransactionItem) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockTransactionItemRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateTransactionItem_Success(t *testing.T) {
	mockRepo := new(MockTransactionItemRepository)
	mockProductRepo := new(MockProductRepository)
	service := NewTransactionItemService(mockRepo, mockProductRepo)

	itemInput := &models.TransactionItem{TransactionID: 1, ProductID: 1, Quantity: 2, Price: 30000}
	product := &models.Product{Name: "Kopi", Stock: 10}

	mockProductRepo.On("FindByID", uint(1)).Return(product, nil)
	mockProductRepo.On("Update", mock.Anything).Return(nil)
	mockRepo.On("Create", itemInput).Return(nil)

	err := service.CreateTransactionItem(itemInput)

	assert.NoError(t, err)
	assert.Equal(t, 8, product.Stock)
	mockRepo.AssertExpectations(t)
	mockProductRepo.AssertExpectations(t)
}

func TestGetAllTransactionItems_Success(t *testing.T) {
	mockRepo := new(MockTransactionItemRepository)
	mockProductRepo := new(MockProductRepository)
	service := NewTransactionItemService(mockRepo, mockProductRepo)

	expectedItems := []models.TransactionItem{
		{TransactionID: 1, Quantity: 2},
		{TransactionID: 2, Quantity: 1},
	}

	mockRepo.On("FindAll").Return(expectedItems, nil)

	result, err := service.GetAllTransactionItems()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, uint(1), result[0].TransactionID)
	mockRepo.AssertExpectations(t)
}

func TestGetTransactionItemByID_NotFound(t *testing.T) {
	mockRepo := new(MockTransactionItemRepository)
	mockProductRepo := new(MockProductRepository)
	service := NewTransactionItemService(mockRepo, mockProductRepo)

	expectedErr := errors.New("record not found")
	mockRepo.On("FindByID", uint(99)).Return(nil, expectedErr)

	result, err := service.GetTransactionItemById(99)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
	mockRepo.AssertExpectations(t)
}

func TestGetTransactionItemsByTransactionID_Success(t *testing.T) {
	mockRepo := new(MockTransactionItemRepository)
	mockProductRepo := new(MockProductRepository)
	service := NewTransactionItemService(mockRepo, mockProductRepo)

	expectedItems := []models.TransactionItem{
		{TransactionID: 1, ProductID: 2},
	}

	mockRepo.On("FindByTransactionID", uint(1)).Return(expectedItems, nil)

	result, err := service.GetTransactionItemsByTransactionId(1)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, uint(1), result[0].TransactionID)
	mockRepo.AssertExpectations(t)
}
