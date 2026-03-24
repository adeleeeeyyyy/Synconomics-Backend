package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type replyRepository struct {
	db *gorm.DB
}

func NewReplyRepository(db *gorm.DB) ReplyRepository {
	return &replyRepository{db}
}

func (r *replyRepository) Create(reply *models.Reply) error {
	return r.db.Create(reply).Error
}

func (r *replyRepository) FindAll() ([]models.Reply, error) {
	var replies []models.Reply
	err := r.db.Preload("User").Preload("Thread").Find(&replies).Error
	return replies, err
}

func (r *replyRepository) FindByID(id uint) (*models.Reply, error) {
	var reply models.Reply
	err := r.db.Preload("User").Preload("Thread").First(&reply, id).Error
	return &reply, err
}

func (r *replyRepository) FindByThreadID(threadID uint) ([]models.Reply, error) {
	var replies []models.Reply
	err := r.db.Preload("User").Where("thread_id = ?", threadID).Find(&replies).Error
	return replies, err
}

func (r *replyRepository) Update(reply *models.Reply) error {
	return r.db.Save(reply).Error
}

func (r *replyRepository) Delete(id uint) error {
	return r.db.Delete(&models.Reply{}, id).Error
}
