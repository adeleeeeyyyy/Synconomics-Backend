package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
)

type replyService struct {
	repo repositories.ReplyRepository
}

func NewReplyService(repo repositories.ReplyRepository) ReplyService {
	return &replyService{repo}
}

func (s *replyService) CreateReply(reply *models.Reply) error {
	return s.repo.Create(reply)
}

func (s *replyService) GetAllReplies() ([]models.Reply, error) {
	return s.repo.FindAll()
}

func (s *replyService) GetReplyById(id uint) (*models.Reply, error) {
	return s.repo.FindByID(id)
}

func (s *replyService) GetRepliesByThreadId(threadID uint) ([]models.Reply, error) {
	return s.repo.FindByThreadID(threadID)
}

func (s *replyService) UpdateReply(reply *models.Reply) error {
	return s.repo.Update(reply)
}

func (s *replyService) DeleteReply(id uint) error {
	return s.repo.Delete(id)
}
