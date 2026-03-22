package services

import (
	"Synconomics/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockExpenseRepository struct {
	mock.Mock
}

func (m *MockExpenseRepository) Create(expense *models.Expense) error {
	args := m.Called(expense)
	return args.Error(0)
}

func (m *MockExpenseRepository) FindAll() ([]models.Expense, error) {
	args := m.Called()
	return args.Get(0).([]models.Expense), args.Error(1)
}

func (m *MockExpenseRepository) FindByID(id uint) (*models.Expense, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Expense), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockExpenseRepository) FindByBusinessID(businessID uint) ([]models.Expense, error) {
	args := m.Called(businessID)
	return args.Get(0).([]models.Expense), args.Error(1)
}

func (m *MockExpenseRepository) Update(expense *models.Expense) error {
	args := m.Called(expense)
	return args.Error(0)
}

func (m *MockExpenseRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateExpense_Success(t *testing.T) {
	mockRepo := new(MockExpenseRepository)
	service := NewExpenseService(mockRepo)

	expenseInput := &models.Expense{Amount: 50000, Category: "Operasional"}

	mockRepo.On("Create", expenseInput).Return(nil)

	err := service.CreateExpense(expenseInput)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllExpenses_Success(t *testing.T) {
	mockRepo := new(MockExpenseRepository)
	service := NewExpenseService(mockRepo)

	expectedExpenses := []models.Expense{
		{Amount: 50000, Category: "Operasional"},
		{Amount: 20000, Category: "Bahan Baku"},
	}

	mockRepo.On("FindAll").Return(expectedExpenses, nil)

	result, err := service.GetAllExpenses()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, float64(50000), float64(result[0].Amount))
	mockRepo.AssertExpectations(t)
}

func TestGetExpenseByID_NotFound(t *testing.T) {
	mockRepo := new(MockExpenseRepository)
	service := NewExpenseService(mockRepo)

	expectedErr := errors.New("record not found")
	mockRepo.On("FindByID", uint(99)).Return(nil, expectedErr)

	result, err := service.GetExpenseById(99)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
	mockRepo.AssertExpectations(t)
}

func TestGetExpenseByBusinessID_Success(t *testing.T) {
	mockRepo := new(MockExpenseRepository)
	service := NewExpenseService(mockRepo)

	expectedExpenses := []models.Expense{
		{BusinessID: 1, Amount: 10000},
	}

	mockRepo.On("FindByBusinessID", uint(1)).Return(expectedExpenses, nil)

	result, err := service.GetExpensesByBusinessId(1)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, uint(1), result[0].BusinessID)
	mockRepo.AssertExpectations(t)
}
