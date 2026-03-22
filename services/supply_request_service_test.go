package services

import (
	"Synconomics/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSupplyRequestRepository struct {
	mock.Mock
}

func (m *MockSupplyRequestRepository) Create(req *models.SupplyRequest) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockSupplyRequestRepository) FindAll() ([]models.SupplyRequest, error) {
	args := m.Called()
	return args.Get(0).([]models.SupplyRequest), args.Error(1)
}

func (m *MockSupplyRequestRepository) FindByID(id uint) (*models.SupplyRequest, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.SupplyRequest), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockSupplyRequestRepository) FindByBusinessID(businessID uint) ([]models.SupplyRequest, error) {
	args := m.Called(businessID)
	return args.Get(0).([]models.SupplyRequest), args.Error(1)
}

func (m *MockSupplyRequestRepository) Update(req *models.SupplyRequest) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockSupplyRequestRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateSupplyRequest_Success(t *testing.T) {
	mockRepo := new(MockSupplyRequestRepository)
	service := NewSupplyRequestService(mockRepo)

	reqInput := &models.SupplyRequest{ProductName: "Gula", Quantity: 10, Status: "open"}

	mockRepo.On("Create", reqInput).Return(nil)

	err := service.CreateSupplyRequest(reqInput)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllSupplyRequests_Success(t *testing.T) {
	mockRepo := new(MockSupplyRequestRepository)
	service := NewSupplyRequestService(mockRepo)

	expectedReqs := []models.SupplyRequest{
		{ProductName: "Gula", Quantity: 10},
		{ProductName: "Kopi", Quantity: 5},
	}

	mockRepo.On("FindAll").Return(expectedReqs, nil)

	result, err := service.GetAllSupplyRequests()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Gula", result[0].ProductName)
	mockRepo.AssertExpectations(t)
}

func TestGetSupplyRequestByID_NotFound(t *testing.T) {
	mockRepo := new(MockSupplyRequestRepository)
	service := NewSupplyRequestService(mockRepo)

	expectedErr := errors.New("record not found")
	mockRepo.On("FindByID", uint(99)).Return(nil, expectedErr)

	result, err := service.GetSupplyRequestById(99)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
	mockRepo.AssertExpectations(t)
}

func TestGetSupplyRequestsByBusinessID_Success(t *testing.T) {
	mockRepo := new(MockSupplyRequestRepository)
	service := NewSupplyRequestService(mockRepo)

	expectedReqs := []models.SupplyRequest{
		{BusinessID: 2, ProductName: "Tepung", Quantity: 20},
	}

	mockRepo.On("FindByBusinessID", uint(2)).Return(expectedReqs, nil)

	result, err := service.GetSupplyRequestsByBusinessId(2)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, uint(2), result[0].BusinessID)
	mockRepo.AssertExpectations(t)
}
