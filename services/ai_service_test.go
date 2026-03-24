package services

import (
	"Synconomics/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAIRepository struct {
	mock.Mock
}

func (m *MockAIRepository) CreateSession(session *models.AISession) error {
	args := m.Called(session)
	return args.Error(0)
}

func (m *MockAIRepository) GetSessionByID(id uint) (*models.AISession, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.AISession), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAIRepository) GetSessionsByUserID(userID uint) ([]models.AISession, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.AISession), args.Error(1)
}

func (m *MockAIRepository) SaveMessage(message *models.AIMessage) error {
	args := m.Called(message)
	return args.Error(0)
}

func (m *MockAIRepository) GetMessagesBySessionID(sessionID uint) ([]models.AIMessage, error) {
	args := m.Called(sessionID)
	return args.Get(0).([]models.AIMessage), args.Error(1)
}

func (m *MockAIRepository) SaveResult(result *models.AIResult) error {
	args := m.Called(result)
	return args.Error(0)
}

func (m *MockAIRepository) GetResultBySessionID(sessionID uint) (*models.AIResult, error) {
	args := m.Called(sessionID)
	if args.Get(0) != nil {
		return args.Get(0).(*models.AIResult), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAIRepository) GetLatestSession(userID, businessID uint, sessionType models.AISessionType) (*models.AISession, error) {
	args := m.Called(userID, businessID, sessionType)
	if args.Get(0) != nil {
		return args.Get(0).(*models.AISession), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestNewAIService(t *testing.T) {
	mockRepo := new(MockAIRepository)
	service := NewAIService(mockRepo)

	assert.NotNil(t, service)
}

func TestCreateSession(t *testing.T) {
	mockRepo := new(MockAIRepository)
	service := NewAIService(mockRepo)

	session := &models.AISession{UserID: 1, BusinessID: 1, Type: models.IdeaGeneration}
	mockRepo.On("CreateSession", session).Return(nil)

	result, err := service.CreateSession(1, 1, string(models.IdeaGeneration))

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}
