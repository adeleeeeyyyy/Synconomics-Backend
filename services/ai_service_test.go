package services

import (
	"Synconomics/models"
	"testing"
	"time"

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

type MockTransactionRepo struct{ mock.Mock }

func (m *MockTransactionRepo) Create(tx *models.Transaction) error { return nil }
func (m *MockTransactionRepo) FindAll() ([]models.Transaction, error) { return nil, nil }
func (m *MockTransactionRepo) FindByID(id uint) (*models.Transaction, error) { return nil, nil }
func (m *MockTransactionRepo) Update(tx *models.Transaction) error { return nil }
func (m *MockTransactionRepo) Delete(id uint) error { return nil }
func (m *MockTransactionRepo) FindByBusinessID(id uint) ([]models.Transaction, error) { return nil, nil }
func (m *MockTransactionRepo) FindByBusinessIDAndDateRange(id uint, s, e time.Time) ([]models.Transaction, error) {
	return nil, nil
}

type MockExpenseRepo struct{ mock.Mock }

func (m *MockExpenseRepo) Create(e *models.Expense) error { return nil }
func (m *MockExpenseRepo) FindAll() ([]models.Expense, error) { return nil, nil }
func (m *MockExpenseRepo) FindByID(id uint) (*models.Expense, error) { return nil, nil }
func (m *MockExpenseRepo) FindByBusinessID(id uint) ([]models.Expense, error) { return nil, nil }
func (m *MockExpenseRepo) Update(e *models.Expense) error { return nil }
func (m *MockExpenseRepo) Delete(id uint) error { return nil }
func (m *MockExpenseRepo) FindByBusinessIDAndDateRange(id uint, s, e time.Time) ([]models.Expense, error) {
	return nil, nil
}

type MockBusinessRepo struct{ mock.Mock }

func (m *MockBusinessRepo) Create(b *models.Business) error { return nil }
func (m *MockBusinessRepo) FindAll() ([]models.Business, error) { return nil, nil }
func (m *MockBusinessRepo) FindByID(id uint) (*models.Business, error) { return nil, nil }
func (m *MockBusinessRepo) Update(b *models.Business) error { return nil }
func (m *MockBusinessRepo) Delete(id uint) error { return nil }
func (m *MockBusinessRepo) FindByUserID(id uint) ([]models.Business, error) { return nil, nil }

type MockProductRepo struct{ mock.Mock }

func (m *MockProductRepo) Create(p *models.Product) error { return nil }
func (m *MockProductRepo) FindAll() ([]models.Product, error) { return nil, nil }
func (m *MockProductRepo) FindByID(id uint) (*models.Product, error) { return nil, nil }
func (m *MockProductRepo) Update(p *models.Product) error { return nil }
func (m *MockProductRepo) Delete(id uint) error { return nil }
func (m *MockProductRepo) FindByBusinessID(id uint) ([]models.Product, error) { return nil, nil }

func TestNewAIService(t *testing.T) {
	mockRepo := new(MockAIRepository)
	service := NewAIService(mockRepo, &MockTransactionRepo{}, &MockExpenseRepo{}, &MockBusinessRepo{}, &MockProductRepo{})

	assert.NotNil(t, service)
}

func TestCreateSession(t *testing.T) {
	mockRepo := new(MockAIRepository)
	service := NewAIService(mockRepo, &MockTransactionRepo{}, &MockExpenseRepo{}, &MockBusinessRepo{}, &MockProductRepo{})

	session := &models.AISession{UserID: 1, BusinessID: 1, Type: models.IdeaGeneration}
	mockRepo.On("CreateSession", session).Return(nil)

	result, err := service.CreateSession(1, 1, string(models.IdeaGeneration))

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}
