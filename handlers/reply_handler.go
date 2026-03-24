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

type ReplyHandler struct {
	service services.ReplyService
}

func NewReplyHandler(service services.ReplyService) *ReplyHandler {
	return &ReplyHandler{service}
}

// CreateReply
// @Summary Membuat Reply baru
// @Description Menambahkan data balasan pada thread
// @Tags replies
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateReplyReq true "Reply Body"
// @Success 201 {object} helpers.Response{data=models.Reply}
// @Router /replies [post]
func (h *ReplyHandler) CreateReply(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var req dto.CreateReplyReq
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	var reply models.Reply
	if err := copier.Copy(&reply, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}
	reply.UserID = userID

	if err := h.service.CreateReply(&reply); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp dto.ReplyResponse
	copier.Copy(&resp, &reply)

	return helpers.SuccessResponse(c, fiber.StatusCreated, "reply created", resp)
}

// GetRepliesByThread
// @Summary Daftar balasan di Thread
// @Description Mengambil semua data balasan milik satu thread
// @Tags replies
// @Produce json
// @Security BearerAuth
// @Param threadId path int true "Thread ID"
// @Success 200 {object} helpers.Response{data=[]dto.ReplyResponse}
// @Router /replies/thread/{threadId} [get]
func (h *ReplyHandler) GetRepliesByThread(c *fiber.Ctx) error {
	threadId, err := strconv.ParseUint(c.Params("threadId"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid thread id")
	}

	replies, err := h.service.GetRepliesByThreadId(uint(threadId))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp []dto.ReplyResponse
	copier.Copy(&resp, &replies)

	return helpers.SuccessResponse(c, fiber.StatusOK, "replies fetched", resp)
}

// UpdateReply
// @Summary Memperbarui Reply
// @Description Mengubah isi balasan
// @Tags replies
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Reply ID"
// @Param request body dto.UpdateReplyReq true "Update Request Body"
// @Success 200 {object} helpers.Response{data=dto.ReplyResponse}
// @Router /replies/{id} [put]
func (h *ReplyHandler) UpdateReply(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	reply, err := h.service.GetReplyById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "reply not found")
	}

	var req dto.UpdateReplyReq
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := copier.Copy(reply, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.UpdateReply(reply); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp dto.ReplyResponse
	copier.Copy(&resp, reply)

	return helpers.SuccessResponse(c, fiber.StatusOK, "reply updated", resp)
}

// DeleteReply
// @Summary Menghapus Reply
// @Description Menghapus balasan
// @Tags replies
// @Produce json
// @Security BearerAuth
// @Param id path int true "Reply ID"
// @Success 200 {object} helpers.Response
// @Router /replies/{id} [delete]
func (h *ReplyHandler) DeleteReply(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.DeleteReply(uint(id)); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "reply deleted", nil)
}
