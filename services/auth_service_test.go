package services

import (
	"Synconomics/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindByGoogleID(googleID string) (*models.User, error) {
	args := m.Called(googleID)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindByID(id uint) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) Create(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Update(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestGetProfile_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewAuthServices(mockRepo)

	mockRepo.On("FindByID", uint(99)).Return(nil, errors.New("user not found"))

	user, token, err := service.GetProfile(99)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Empty(t, token)
	assert.Equal(t, "user not found", err.Error())
	mockRepo.AssertExpectations(t)
}

// Minimal placeholder tests for Register and Login since they involve password hashing and JWT generation.
func TestLogin_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewAuthServices(mockRepo)

	mockRepo.On("FindByEmail", "unknown@example.com").Return(nil, errors.New("record not found"))

	user, token, err := service.Login("unknown@example.com", "pass123")
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Empty(t, token)
	assert.Equal(t, "invalid credential", err.Error())
	mockRepo.AssertExpectations(t)
}
