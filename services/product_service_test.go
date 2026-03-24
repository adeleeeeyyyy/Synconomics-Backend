package services

import (
	"Synconomics/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Create(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) FindAll() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepository) FindByID(id uint) (*models.Product, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockProductRepository) Update(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockProductRepository) FindByBusinessID(businessID uint) ([]models.Product, error) {
	args := m.Called(businessID)
	return args.Get(0).([]models.Product), args.Error(1)
}

func TestCreateProduct_Success(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := NewProductService(mockRepo)

	productInput := &models.Product{Name: "Kopi Susu", Price: 15000}

	mockRepo.On("Create", productInput).Return(nil)

	err := service.CreateProduct(productInput)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllProducts_Success(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := NewProductService(mockRepo)

	expectedProducts := []models.Product{
		{Name: "Kopi Susu", Price: 15000},
		{Name: "Roti Bakar", Price: 20000},
	}

	mockRepo.On("FindAll").Return(expectedProducts, nil)

	result, err := service.GetAllProducts()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Kopi Susu", result[0].Name)
	mockRepo.AssertExpectations(t)
}

func TestGetProductByID_NotFound(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := NewProductService(mockRepo)

	expectedErr := errors.New("record not found")
	mockRepo.On("FindByID", uint(99)).Return(nil, expectedErr)

	result, err := service.GetProductById(99)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
	mockRepo.AssertExpectations(t)
}

func TestGetProductsByBusinessId_Success(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := NewProductService(mockRepo)

	expectedProducts := []models.Product{
		{Name: "Kopi Susu", BusinessID: 1},
	}

	mockRepo.On("FindByBusinessID", uint(1)).Return(expectedProducts, nil)

	result, err := service.GetProductsByBusinessId(1)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, uint(1), result[0].BusinessID)
	mockRepo.AssertExpectations(t)
}