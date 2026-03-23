package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type aiService struct {
	repo repositories.AIRepository
}

func NewAIService(repo repositories.AIRepository) AIService {
	return &aiService{repo}
}

func (s *aiService) CreateSession(userID, businessID uint, sessionType string) (*models.AISession, error) {
	session := &models.AISession{
		UserID:     userID,
		BusinessID: businessID,
		Type:       models.AISessionType(sessionType),
	}
	if err := s.repo.CreateSession(session); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *aiService) Chat(sessionID uint, userMessage string) (*models.AIMessage, error) {
	session, err := s.repo.GetSessionByID(sessionID)
	if err != nil {
		return nil, err
	}

	userMsg := &models.AIMessage{
		AISessionID: sessionID,
		Role:        "user",
		Content:     userMessage,
	}
	if err := s.repo.SaveMessage(userMsg); err != nil {
		return nil, err
	}

	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash-latest")

	switch session.Type {
	case models.IdeaGeneration:
		model.SystemInstruction = &genai.Content{Parts: []genai.Part{genai.Text("You are an expert AI business consultant specializing in Idea Generation.")}}
	case models.Validation:
		model.SystemInstruction = &genai.Content{Parts: []genai.Part{genai.Text("You are an expert AI business consultant specializing in Business Validation.")}}
	case models.Strategy:
		model.SystemInstruction = &genai.Content{Parts: []genai.Part{genai.Text("You are an expert AI business consultant specializing in Business Strategy.")}}
	}

	cs := model.StartChat()
	
	messages, err := s.repo.GetMessagesBySessionID(sessionID)
	if err == nil {
		var history []*genai.Content
		for _, msg := range messages {
			if msg.ID == userMsg.ID {
				continue
			}
			history = append(history, &genai.Content{
				Role:  msg.Role,
				Parts: []genai.Part{genai.Text(msg.Content)},
			})
		}
		cs.History = history
	}

	resp, err := cs.SendMessage(ctx, genai.Text(userMessage))
	if err != nil {
		return nil, fmt.Errorf("gemini sdk error: %v", err)
	}

	var modelResponseText string
	for _, part := range resp.Candidates[0].Content.Parts {
		if textPart, ok := part.(genai.Text); ok {
			modelResponseText += string(textPart)
		}
	}

	modelMsg := &models.AIMessage{
		AISessionID: sessionID,
		Role:        "model",
		Content:     modelResponseText,
	}
	if err := s.repo.SaveMessage(modelMsg); err != nil {
		return nil, err
	}

	return modelMsg, nil
}

func (s *aiService) FinalizeSessionResult(sessionID uint) (*models.AIResult, error) {
	messages, err := s.repo.GetMessagesBySessionID(sessionID)
	if err != nil || len(messages) == 0 {
		return nil, fmt.Errorf("sesi kosong atau tidak ditemukan")
	}

	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash-latest")
	model.ResponseMIMEType = "application/json"

	var fullConversation string
	for _, m := range messages {
		fullConversation += fmt.Sprintf("[%s]: %s\n", m.Role, m.Content)
	}

	prompt := fmt.Sprintf("Analyze the conversation and provide a JSON result. Conversation:\n%s", fullConversation)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	var jsonOutput string
	for _, part := range resp.Candidates[0].Content.Parts {
		if textPart, ok := part.(genai.Text); ok {
			jsonOutput += string(textPart)
		}
	}

	result := &models.AIResult{
		AISessionID: sessionID,
		Data:        jsonOutput,
	}
	
	session, _ := s.repo.GetSessionByID(sessionID)
	if session != nil {
		result.ResultType = string(session.Type)
	}

	if err := s.repo.SaveResult(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *aiService) GetSessionMessages(sessionID uint) ([]models.AIMessage, error) {
	return s.repo.GetMessagesBySessionID(sessionID)
}

func (s *aiService) GetSessionResult(sessionID uint) (*models.AIResult, error) {
	return s.repo.GetResultBySessionID(sessionID)
}

func (s *aiService) ChatByRole(userID, businessID uint, sessionType string, message string) (*models.AIMessage, error) {
	session, err := s.repo.GetLatestSession(userID, businessID, models.AISessionType(sessionType))
	if err != nil {
		// Create new session if not found
		session, err = s.CreateSession(userID, businessID, sessionType)
		if err != nil {
			return nil, err
		}
	}

	return s.Chat(session.ID, message)
}
