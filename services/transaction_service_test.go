package services

import (
	"Synconomics/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(transaction *models.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) FindAll() ([]models.Transaction, error) {
	args := m.Called()
	return args.Get(0).([]models.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) FindByID(id uint) (*models.Transaction, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Transaction), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTransactionRepository) Update(transaction *models.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTransactionRepository) FindByBusinessID(businessID uint) ([]models.Transaction, error) {
	args := m.Called(businessID)
	return args.Get(0).([]models.Transaction), args.Error(1)
}

func TestCreateTransaction_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)

	txInput := &models.Transaction{TotalAmount: 150000, BusinessID: 1}

	mockRepo.On("Create", txInput).Return(nil)

	err := service.CreateTransaction(txInput)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllTransactions_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)

	expectedTxs := []models.Transaction{
		{TotalAmount: 150000},
		{TotalAmount: 50000},
	}

	mockRepo.On("FindAll").Return(expectedTxs, nil)

	result, err := service.GetAllTransactions()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, float64(150000), result[0].TotalAmount)
	mockRepo.AssertExpectations(t)
}

func TestGetTransactionByID_NotFound(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)

	expectedErr := errors.New("record not found")
	mockRepo.On("FindByID", uint(99)).Return(nil, expectedErr)

	result, err := service.GetTransactionById(99)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
	mockRepo.AssertExpectations(t)
}

func TestGetTransactionsByBusinessId_Success(t *testing.T) {
	mockRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockRepo)

	expectedTxs := []models.Transaction{
		{TotalAmount: 150000, BusinessID: 1},
	}

	mockRepo.On("FindByBusinessID", uint(1)).Return(expectedTxs, nil)

	result, err := service.GetTransactionsByBusinessId(1)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, float64(150000), result[0].TotalAmount)
	mockRepo.AssertExpectations(t)
}
