package handlers

import (
	"Synconomics/dto"
	"Synconomics/models"
	"Synconomics/pkg/helpers"
	"Synconomics/services"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type TransactionHandler struct {
	service services.TransactionService
}

func NewTransactionHandler(service services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service}
}

// CreateTransaction
// @Summary Membuat transaksi baru
// @Description Menambahkan data transaksi baru
// @Tags transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateTransactionRequest true "Transaction Request"
// @Success 201 {object} helpers.Response{data=models.Transaction}
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /transactions [post]
func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	var req dto.CreateTransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	var newTransaction models.Transaction
	if err := copier.Copy(&newTransaction, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	newTransaction.TransactionDate = time.Now()

	if err := h.service.CreateTransaction(&newTransaction); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp dto.TransactionResponse
	copier.Copy(&resp, &newTransaction)

	return helpers.SuccessResponse(c, fiber.StatusCreated, "transaction created", resp)
}

// GetTransactions
// @Summary ambil semua daftar transaksi
// @Description mengambil semua data transaksi yang ada
// @Tags transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]dto.TransactionResponse} "Sukses mengambil data"
// @Failure 500 {object} helpers.Response "internal server error"
// @Router /transactions [get]
func (h *TransactionHandler) GetTransactions(c *fiber.Ctx) error {
	transactions, err := h.service.GetAllTransactions()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp []dto.TransactionResponse
	copier.Copy(&resp, &transactions)

	return helpers.SuccessResponse(c, fiber.StatusOK, "transactions fetched", resp)
}

// GetTransaction
// @Summary Mendapatkan detail transaksi
// @Description Mengambil informasi lengkap terkait satu transaksi menggunakan ID
// @Tags transactions
// @Produce json
// @Security BearerAuth
// @Param id path int true "Transaction ID"
// @Success 200 {object} helpers.Response{data=dto.TransactionResponse}
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetTransaction(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	transaction, err := h.service.GetTransactionById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "transaction not found")
	}

	var resp dto.TransactionResponse
	copier.Copy(&resp, transaction)

	return helpers.SuccessResponse(c, fiber.StatusOK, "transaction fetched", resp)
}

// UpdateTransaction
// @Summary Memperbarui data transaksi
// @Description Mengubah kolom dari sebuah entitas transaksi
// @Tags transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Transaction ID"
// @Param request body dto.UpdateTransactionRequest true "Transaction Update Request"
// @Success 200 {object} helpers.Response{data=dto.TransactionResponse}
// @Router /transactions/{id} [put]
// GetTransactionsByBusinessId
// @Summary Mendapatkan detail transaksi berdasarkan Business ID
// @Description Mengambil kumpulan transaksi menggunakan Business ID
// @Tags transactions
// @Produce json
// @Security BearerAuth
// @Param id path int true "Business ID"
// @Success 200 {object} helpers.Response{data=[]dto.TransactionResponse}
// @Router /transactions/business/{id} [get]
func (h *TransactionHandler) GetTransactionsByBusinessId(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid business id")
	}

	transactions, err := h.service.GetTransactionsByBusinessId(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "transactions not found")
	}

	var resp []dto.TransactionResponse
	copier.Copy(&resp, &transactions)

	return helpers.SuccessResponse(c, fiber.StatusOK, "transactions fetched", resp)
}

func (h *TransactionHandler) UpdateTransaction(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	existingTransaction, err := h.service.GetTransactionById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "transaction not found")
	}

	var req dto.UpdateTransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := copier.Copy(existingTransaction, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.UpdateTransaction(existingTransaction); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp dto.TransactionResponse
	copier.Copy(&resp, existingTransaction)

	return helpers.SuccessResponse(c, fiber.StatusOK, "transaction updated", resp)
}

// DeleteTransaction
// @Summary Menghapus transaksi
// @Description Menghapus transaksi dari database
// @Tags transactions
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Transaction ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /transactions/{id} [delete]
func (h *TransactionHandler) DeleteTransaction(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.DeleteTransaction(uint(id)); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "transaction deleted", nil)
}
