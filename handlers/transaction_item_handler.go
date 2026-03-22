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

type TransactionItemHandler struct {
	service services.TransactionItemService
}

func NewTransactionItemHandler(service services.TransactionItemService) *TransactionItemHandler {
	return &TransactionItemHandler{service}
}

// CreateTransactionItem
// @Summary Membuat transaction_item baru
// @Description Menambahkan data entitas transaction_item baru
// @Tags transaction_items
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateTransactionItemRequest true "Transaction Item Request"
// @Success 201 {object} helpers.Response{data=models.TransactionItem}
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /transaction_items [post]
func (h *TransactionItemHandler) CreateTransactionItem(c *fiber.Ctx) error {
	var req dto.CreateTransactionItemRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	var newItem models.TransactionItem
	if err := copier.Copy(&newItem, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.CreateTransactionItem(&newItem); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusCreated, "transaction item created", newItem)
}

// GetTransactionItems
// @Summary ambil semua daftar transaction_items
// @Description mengambil semua data transaction_items yang ada
// @Tags transaction_items
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]models.TransactionItem}
// @Failure 500 {object} helpers.Response
// @Router /transaction_items [get]
func (h *TransactionItemHandler) GetTransactionItems(c *fiber.Ctx) error {
	items, err := h.service.GetAllTransactionItems()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "transaction items fetched", items)
}

// GetTransactionItem
// @Summary Mendapatkan detail transaction_item
// @Description Mengambil informasi lengkap terkait satu transaction_item menggunakan ID
// @Tags transaction_items
// @Produce json
// @Security BearerAuth
// @Param id path int true "Transaction Item ID"
// @Success 200 {object} helpers.Response{data=models.TransactionItem}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /transaction_items/{id} [get]
func (h *TransactionItemHandler) GetTransactionItem(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	item, err := h.service.GetTransactionItemById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "transaction item not found")
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "transaction item fetched", item)
}

// GetTransactionItemsByTransactionId
// @Summary Mendapatkan detail transaction_item berdasarkan Transaction ID
// @Description Mengambil informasi lengkap terkait kumpulan transaction_item menggunakan Transaction ID
// @Tags transaction_items
// @Produce json
// @Security BearerAuth
// @Param id path int true "Transaction ID"
// @Success 200 {object} helpers.Response{data=[]models.TransactionItem}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /transaction_items/transaction/{id} [get]
func (h *TransactionItemHandler) GetTransactionItemsByTransactionId(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid transaction id")
	}

	items, err := h.service.GetTransactionItemsByTransactionId(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "transaction items not found")
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "transaction items fetched", items)
}

// UpdateTransactionItem
// @Summary Memperbarui data transaction_item
// @Description Mengubah kolom kuantitas dan harga dari sebuah entitas transaction_item
// @Tags transaction_items
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Transaction Item ID"
// @Param request body dto.UpdateTransactionItemRequest true "Transaction Item Update Request"
// @Success 200 {object} helpers.Response{data=models.TransactionItem}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /transaction_items/{id} [put]
func (h *TransactionItemHandler) UpdateTransactionItem(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	existingItem, err := h.service.GetTransactionItemById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "transaction item not found")
	}

	var req dto.UpdateTransactionItemRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := copier.Copy(existingItem, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.UpdateTransactionItem(existingItem); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "transaction item updated", existingItem)
}

// DeleteTransactionItem
// @Summary Menghapus transaction_item
// @Description Menghapus transaction_item dari database
// @Tags transaction_items
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Transaction Item ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /transaction_items/{id} [delete]
func (h *TransactionItemHandler) DeleteTransactionItem(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.DeleteTransactionItem(uint(id)); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "transaction item deleted", nil)
}
