package services

import (
	"Synconomics/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBusinessRepository simulates the database repository for Business
type MockBusinessRepository struct {
	mock.Mock
}

func (m *MockBusinessRepository) Create(business *models.Business) error {
	args := m.Called(business)
	return args.Error(0)
}

func (m *MockBusinessRepository) FindAll() ([]models.Business, error) {
	args := m.Called()
	return args.Get(0).([]models.Business), args.Error(1)
}

func (m *MockBusinessRepository) FindByID(id uint) (*models.Business, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Business), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockBusinessRepository) Update(business *models.Business) error {
	args := m.Called(business)
	return args.Error(0)
}

func (m *MockBusinessRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBusinessRepository) FindByUserID(userID uint) ([]models.Business, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.Business), args.Error(1)
}

// ---------------------------------------------
// Test Functions for BusinessService
// ---------------------------------------------

func TestCreateBusiness_Success(t *testing.T) {
	mockRepo := new(MockBusinessRepository)
	service := NewBusinessService(mockRepo)

	businessInput := &models.Business{Name: "Toko Sinar Jaya", Category: "Retail"}

	mockRepo.On("Create", businessInput).Return(nil)

	err := service.CreateBusiness(businessInput)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllBusinesses_Success(t *testing.T) {
	mockRepo := new(MockBusinessRepository)
	service := NewBusinessService(mockRepo)

	expectedBusinesses := []models.Business{
		{Name: "Toko Sinar Jaya", Category: "Retail"},
		{Name: "Warung Bu Anna", Category: "F&B"},
	}

	mockRepo.On("FindAll").Return(expectedBusinesses, nil)

	result, err := service.GetAllBusinesses()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Toko Sinar Jaya", result[0].Name)
	mockRepo.AssertExpectations(t)
}

func TestGetBusinessById_NotFound(t *testing.T) {
	mockRepo := new(MockBusinessRepository)
	service := NewBusinessService(mockRepo)

	expectedErr := errors.New("business record not found")
	mockRepo.On("FindByID", uint(99)).Return(nil, expectedErr)

	result, err := service.GetBusinessById(99)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateBusiness_Success(t *testing.T) {
	mockRepo := new(MockBusinessRepository)
	service := NewBusinessService(mockRepo)

	businessToUpdate := &models.Business{Name: "Toko Sinar Jaya Baru", Category: "Retail", Address: "Jl. ABC"}

	mockRepo.On("Update", businessToUpdate).Return(nil)

	err := service.UpdateBusiness(businessToUpdate)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteBusiness_Success(t *testing.T) {
	mockRepo := new(MockBusinessRepository)
	service := NewBusinessService(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := service.DeleteBusiness(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetBusinessesByUserId_Success(t *testing.T) {
	mockRepo := new(MockBusinessRepository)
	service := NewBusinessService(mockRepo)

	expectedBusinesses := []models.Business{
		{Name: "Toko Sinar Jaya", UserID: 1},
	}

	mockRepo.On("FindByUserID", uint(1)).Return(expectedBusinesses, nil)

	result, err := service.GetBusinessesByUserId(1)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, uint(1), result[0].UserID)
	mockRepo.AssertExpectations(t)
}
