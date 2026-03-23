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

type ThreadHandler struct {
	service services.ThreadService
}

func NewThreadHandler(service services.ThreadService) *ThreadHandler {
	return &ThreadHandler{service}
}

// CreateThread
// @Summary Membuat Thread baru
// @Description Menambahkan data diskusi/thread baru
// @Tags threads
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateThreadReq true "Thread Body"
// @Success 201 {object} helpers.Response{data=models.Thread}
// @Router /threads [post]
func (h *ThreadHandler) CreateThread(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var req dto.CreateThreadReq
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	var thread models.Thread
	if err := copier.Copy(&thread, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}
	thread.UserID = userID

	if err := h.service.CreateThread(&thread); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusCreated, "thread created", thread)
}

// GetThreads
// @Summary Daftar seluruh Thread
// @Description Mengambil semua data thread
// @Tags threads
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]models.Thread}
// @Router /threads [get]
func (h *ThreadHandler) GetThreads(c *fiber.Ctx) error {
	threads, err := h.service.GetAllThreads()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "threads fetched", threads)
}

// GetThread
// @Summary Detail Thread
// @Description Mengambil informasi thread menggunakan ID
// @Tags threads
// @Produce json
// @Security BearerAuth
// @Param id path int true "Thread ID"
// @Success 200 {object} helpers.Response{data=models.Thread}
// @Router /threads/{id} [get]
func (h *ThreadHandler) GetThread(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	thread, err := h.service.GetThreadById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "thread not found")
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "thread fetched", thread)
}

// UpdateThread
// @Summary Memperbarui Thread
// @Description Mengubah kolom pada thread
// @Tags threads
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Thread ID"
// @Param request body dto.UpdateThreadReq true "Update Request Body"
// @Success 200 {object} helpers.Response{data=models.Thread}
// @Router /threads/{id} [put]
func (h *ThreadHandler) UpdateThread(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	thread, err := h.service.GetThreadById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "thread not found")
	}

	var req dto.UpdateThreadReq
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := copier.Copy(thread, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.UpdateThread(thread); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "thread updated", thread)
}

// DeleteThread
// @Summary Menghapus Thread
// @Description Menghapus thread
// @Tags threads
// @Produce json
// @Security BearerAuth
// @Param id path int true "Thread ID"
// @Success 200 {object} helpers.Response
// @Router /threads/{id} [delete]
func (h *ThreadHandler) DeleteThread(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.DeleteThread(uint(id)); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "thread deleted", nil)
}
