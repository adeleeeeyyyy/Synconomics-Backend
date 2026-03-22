package handlers

import (
	"Synconomics/dto"
	"Synconomics/pkg/helpers"
	"Synconomics/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AIHandler struct {
	service services.AIService
}

func NewAIHandler(service services.AIService) *AIHandler {
	return &AIHandler{service}
}

// CreateSession
// @Summary Membuat sesi AI chat baru
// @Description Menyimpan topik dan inisialisasi session AI
// @Tags ai
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateSessionRequest true "Create Session Request"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /ai/sessions [post]
func (h *AIHandler) CreateSession(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return helpers.ErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	var req dto.CreateSessionRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	session, err := h.service.CreateSession(userID, req.BusinessID, req.Type)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusCreated, "ai session created", session)
}

// Chat
// @Summary Mengirim pesan ke Gemini AI
// @Description Membalas pesan prompt user berdasarkan konteks session
// @Tags ai
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Session ID"
// @Param request body dto.ChatRequest true "Chat Request"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /ai/sessions/{id}/chat [post]
func (h *AIHandler) Chat(c *fiber.Ctx) error {
	sessionID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid session id")
	}

	var req dto.ChatRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	reply, err := h.service.Chat(uint(sessionID), req.Message)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "ai message replied", reply)
}

// GetMessages
// @Summary Mengambil riwayat chat
// @Description Mendapatkan seluruh log AIMessage berdasarkan Session ID
// @Tags ai
// @Produce json
// @Security BearerAuth
// @Param id path int true "Session ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /ai/sessions/{id}/messages [get]
func (h *AIHandler) GetMessages(c *fiber.Ctx) error {
	sessionID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid session id")
	}

	messages, err := h.service.GetSessionMessages(uint(sessionID))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "messages fetched", messages)
}

// FinalizeResult
// @Summary Mengkalkulasi resume percakapan chat session
// @Description Menutup sesi dan merangkum prompt di json AIResult
// @Tags ai
// @Produce json
// @Security BearerAuth
// @Param id path int true "Session ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /ai/sessions/{id}/result [post]
func (h *AIHandler) FinalizeResult(c *fiber.Ctx) error {
	sessionID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid session id")
	}

	result, err := h.service.FinalizeSessionResult(uint(sessionID))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "session finalized", result)
}
