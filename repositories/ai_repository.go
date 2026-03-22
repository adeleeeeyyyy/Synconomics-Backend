package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type aiRepository struct {
	db *gorm.DB
}

func NewAIRepository(db *gorm.DB) AIRepository {
	return &aiRepository{db}
}

func (r *aiRepository) CreateSession(session *models.AISession) error {
	return r.db.Create(session).Error
}

func (r *aiRepository) GetSessionByID(id uint) (*models.AISession, error) {
	var session models.AISession
	err := r.db.Preload("Business").Preload("User").First(&session, id).Error
	return &session, err
}

func (r *aiRepository) GetSessionsByUserID(userID uint) ([]models.AISession, error) {
	var sessions []models.AISession
	err := r.db.Preload("Business").Where("user_id = ?", userID).Find(&sessions).Error
	return sessions, err
}

func (r *aiRepository) SaveMessage(message *models.AIMessage) error {
	return r.db.Create(message).Error
}

func (r *aiRepository) GetMessagesBySessionID(sessionID uint) ([]models.AIMessage, error) {
	var messages []models.AIMessage
	err := r.db.Where("ai_session_id = ?", sessionID).Order("created_at asc").Find(&messages).Error
	return messages, err
}

func (r *aiRepository) SaveResult(result *models.AIResult) error {
	return r.db.Create(result).Error
}

func (r *aiRepository) GetResultBySessionID(sessionID uint) (*models.AIResult, error) {
	var result models.AIResult
	err := r.db.Where("ai_session_id = ?", sessionID).First(&result).Error
	return &result, err
}
