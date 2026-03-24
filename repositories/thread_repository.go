package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type threadRepository struct {
	db *gorm.DB
}

func NewThreadRepository(db *gorm.DB) ThreadRepository {
	return &threadRepository{db}
}

func (r *threadRepository) Create(thread *models.Thread) error {
	return r.db.Preload("User").Create(thread).Error
}

func (r *threadRepository) FindAll() ([]models.Thread, error) {
	var threads []models.Thread
	err := r.db.Preload("User").Preload("Replies").Find(&threads).Error
	return threads, err
}

func (r *threadRepository) FindByID(id uint) (*models.Thread, error) {
	var thread models.Thread
	err := r.db.Preload("User").Preload("Replies").Preload("Replies.User").First(&thread, id).Error
	return &thread, err
}

func (r *threadRepository) Update(thread *models.Thread) error {
	return r.db.Save(thread).Error
}

func (r *threadRepository) Delete(id uint) error {
	return r.db.Delete(&models.Thread{}, id).Error
}
