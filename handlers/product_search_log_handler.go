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

type ProductSearchLogHandler struct {
	service services.ProductSearchLogService
}

func NewProductSearchLogHandler(service services.ProductSearchLogService) *ProductSearchLogHandler {
	return &ProductSearchLogHandler{service}
}

// CreateLog
// @Summary Membuat Log Pencarian Produk
// @Description Menyimpan kata kunci pencarian yang dilakukan user
// @Tags product_search_logs
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateProductSearchLogReq true "Log Body"
// @Success 201 {object} helpers.Response{data=models.ProductSearchLog}
// @Router /product-search-logs [post]
func (h *ProductSearchLogHandler) CreateLog(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	var req dto.CreateProductSearchLogReq
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	var log models.ProductSearchLog
	if err := copier.Copy(&log, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}
	log.UserID = userID

	if err := h.service.CreateLog(&log); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusCreated, "search log created", log)
}

// GetLogs
// @Summary Daftar seluruh Log Pencarian
// @Description Mengambil semua data log pencarian produk
// @Tags product_search_logs
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]models.ProductSearchLog}
// @Router /product-search-logs [get]
func (h *ProductSearchLogHandler) GetLogs(c *fiber.Ctx) error {
	logs, err := h.service.GetAllLogs()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "search logs fetched", logs)
}

// GetLog
// @Summary Detail Log Pencarian
// @Description Mengambil informasi log pencarian menggunakan ID
// @Tags product_search_logs
// @Produce json
// @Security BearerAuth
// @Param id path int true "Log ID"
// @Success 200 {object} helpers.Response{data=models.ProductSearchLog}
// @Router /product-search-logs/{id} [get]
func (h *ProductSearchLogHandler) GetLog(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	log, err := h.service.GetLogById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "search log not found")
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "search log fetched", log)
}

// GetUserLogs
// @Summary Daftar Log Pencarian milik User
// @Description Mengambil semua data log pencarian milik user yang sedang login
// @Tags product_search_logs
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]models.ProductSearchLog}
// @Router /product-search-logs/me [get]
func (h *ProductSearchLogHandler) GetUserLogs(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	logs, err := h.service.GetLogsByUserId(userID)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "user search logs fetched", logs)
}
