package services

import (
	"Synconomics/models"
	"errors"
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

func TestCreateSession_Success(t *testing.T) {
	mockRepo := new(MockAIRepository)
	service := NewAIService(mockRepo)

	mockRepo.On("CreateSession", mock.AnythingOfType("*models.AISession")).Return(nil)

	session, err := service.CreateSession(1, 1, "business_analysis")

	assert.NoError(t, err)
	assert.NotNil(t, session)
	assert.Equal(t, uint(1), session.UserID)
	assert.Equal(t, uint(1), session.BusinessID)
	mockRepo.AssertExpectations(t)
}

func TestGetSessionMessages_NotFound(t *testing.T) {
	mockRepo := new(MockAIRepository)
	service := NewAIService(mockRepo)

	mockRepo.On("GetMessagesBySessionID", uint(99)).Return([]models.AIMessage{}, errors.New("record not found"))

	msgs, err := service.GetSessionMessages(99)

	assert.Error(t, err)
	assert.Empty(t, msgs)
	mockRepo.AssertExpectations(t)
}

func TestGetSessionResult_NotFound(t *testing.T) {
	mockRepo := new(MockAIRepository)
	service := NewAIService(mockRepo)

	mockRepo.On("GetResultBySessionID", uint(99)).Return(nil, errors.New("record not found"))

	res, err := service.GetSessionResult(99)

	assert.Error(t, err)
	assert.Nil(t, res)
	mockRepo.AssertExpectations(t)
}
