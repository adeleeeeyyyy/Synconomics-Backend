package handlers

import (
	"Synconomics/dto"
	"Synconomics/models"
	"Synconomics/pkg/helpers"
	"Synconomics/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
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
// @Success 201 {object} helpers.Response{data=dto.AISessionResponse}
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

	var resp dto.AISessionResponse
	copier.Copy(&resp, session)

	return helpers.SuccessResponse(c, fiber.StatusCreated, "ai session created", resp)
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
// @Success 200 {object} helpers.Response{data=dto.AIMessageResponse}
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

	token := c.Get("Authorization")
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	reply, err := h.service.Chat(uint(sessionID), req.Message, token)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp dto.AIMessageResponse
	copier.Copy(&resp, reply)

	return helpers.SuccessResponse(c, fiber.StatusOK, "ai message replied", resp)
}

// GetMessages
// @Summary Mengambil riwayat chat
// @Description Mendapatkan seluruh log AIMessage berdasarkan Session ID
// @Tags ai
// @Produce json
// @Security BearerAuth
// @Param id path int true "Session ID"
// @Success 200 {object} helpers.Response{data=[]dto.AIMessageResponse}
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

	var resp []dto.AIMessageResponse
	copier.Copy(&resp, &messages)

	return helpers.SuccessResponse(c, fiber.StatusOK, "messages fetched", resp)
}

// FinalizeResult
// @Summary Mengkalkulasi resume percakapan chat session
// @Description Menutup sesi dan merangkum prompt di json AIResult
// @Tags ai
// @Produce json
// @Security BearerAuth
// @Param id path int true "Session ID"
// @Success 200 {object} helpers.Response{data=dto.AIResultResponse}
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

	var resp dto.AIResultResponse
	copier.Copy(&resp, result)

	return helpers.SuccessResponse(c, fiber.StatusOK, "session finalized", resp)
}

// ChatIdeaGeneration
// @Summary Chat khusus Idea Generation
// @Description Mengirim pesan ke AI dengan role Idea Generation (otomatis session)
// @Tags ai
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.ChatWithRoleRequest true "Chat Request"
// @Success 200 {object} helpers.Response
// @Router /ai/chat/idea-generation [post]
func (h *AIHandler) ChatIdeaGeneration(c *fiber.Ctx) error {
	return h.chatByRole(c, string(models.IdeaGeneration))
}

// ChatValidation
// @Summary Chat khusus Business Validation
// @Description Mengirim pesan ke AI dengan role Business Validation (otomatis session)
// @Tags ai
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.ChatWithRoleRequest true "Chat Request"
// @Success 200 {object} helpers.Response
// @Router /ai/chat/validation [post]
func (h *AIHandler) ChatValidation(c *fiber.Ctx) error {
	return h.chatByRole(c, string(models.Validation))
}

// ChatStrategy
// @Summary Chat khusus Business Strategy
// @Description Mengirim pesan ke AI dengan role Business Strategy (otomatis session)
// @Tags ai
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.ChatWithRoleRequest true "Chat Request"
// @Success 200 {object} helpers.Response{data=dto.AIMessageResponse}
// @Router /ai/chat/strategy [post]
func (h *AIHandler) ChatStrategy(c *fiber.Ctx) error {
	return h.chatByRole(c, string(models.Strategy))
}

func (h *AIHandler) chatByRole(c *fiber.Ctx, sessionType string) error {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return helpers.ErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	var req dto.ChatWithRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	token := c.Get("Authorization")
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	reply, err := h.service.ChatByRole(userID, req.BusinessID, sessionType, req.Message, token)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp dto.AIMessageResponse
	copier.Copy(&resp, reply)

	return helpers.SuccessResponse(c, fiber.StatusOK, "ai message replied", resp)
}

// AuditBusiness
// @Summary Audit laporan bisnis user via AI
// @Description Menganalisis data transaksi dan pengeluaran 30 hari terakhir untuk memberikan audit finansial.
// @Tags ai
// @Produce json
// @Security BearerAuth
// @Param business_id path int true "Business ID"
// @Success 200 {object} helpers.Response{data=string}
// @Router /ai/audit/{business_id} [post]
func (h *AIHandler) AuditBusiness(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return helpers.ErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	businessID, err := strconv.ParseUint(c.Params("business_id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid business id")
	}

	token := c.Get("Authorization")
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	auditResult, err := h.service.AuditBusinessReport(userID, uint(businessID), token)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "business audit completed", auditResult)
}
