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

type ExpenseHandler struct {
	service services.ExpenseService
}

func NewExpenseHandler(service services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{service}
}

// CreateExpense
// @Summary Membuat expense baru
// @Description Menambahkan data entitas expense baru
// @Tags expenses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateExpenseRequest true "Expense Request"
// @Success 201 {object} helpers.Response{data=models.Expense}
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /expenses [post]
func (h *ExpenseHandler) CreateExpense(c *fiber.Ctx) error {
	var req dto.CreateExpenseRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	var newExpense models.Expense
	if err := copier.Copy(&newExpense, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.CreateExpense(&newExpense); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusCreated, "expense created", newExpense)
}

// GetExpenses
// @Summary ambil semua daftar expenses
// @Description mengambil semua data expenses yang ada
// @Tags expenses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]models.Expense}
// @Failure 500 {object} helpers.Response
// @Router /expenses [get]
func (h *ExpenseHandler) GetExpenses(c *fiber.Ctx) error {
	expenses, err := h.service.GetAllExpenses()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "expenses fetched", expenses)
}

// GetExpense
// @Summary Mendapatkan detail expense
// @Description Mengambil informasi lengkap terkait satu expense menggunakan ID
// @Tags expenses
// @Produce json
// @Security BearerAuth
// @Param id path int true "Expense ID"
// @Success 200 {object} helpers.Response{data=models.Expense}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /expenses/{id} [get]
func (h *ExpenseHandler) GetExpense(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	expense, err := h.service.GetExpenseById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "expense not found")
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "expense fetched", expense)
}

// GetExpensesByBusinessId
// @Summary Mendapatkan detail expense berdasarkan Business ID
// @Description Mengambil kumpulan expense menggunakan Business ID
// @Tags expenses
// @Produce json
// @Security BearerAuth
// @Param id path int true "Business ID"
// @Success 200 {object} helpers.Response{data=[]models.Expense}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /expenses/business/{id} [get]
func (h *ExpenseHandler) GetExpensesByBusinessId(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid business id")
	}

	expenses, err := h.service.GetExpensesByBusinessId(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "expenses not found")
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "expenses fetched", expenses)
}

// UpdateExpense
// @Summary Memperbarui data expense
// @Description Mengubah kolom expense
// @Tags expenses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Expense ID"
// @Param request body dto.UpdateExpenseRequest true "Expense Update Request"
// @Success 200 {object} helpers.Response{data=models.Expense}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /expenses/{id} [put]
func (h *ExpenseHandler) UpdateExpense(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	existingExpense, err := h.service.GetExpenseById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "expense not found")
	}

	var req dto.UpdateExpenseRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := copier.Copy(existingExpense, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.UpdateExpense(existingExpense); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "expense updated", existingExpense)
}

// DeleteExpense
// @Summary Menghapus expense
// @Description Menghapus expense dari database
// @Tags expenses
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Expense ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /expenses/{id} [delete]
func (h *ExpenseHandler) DeleteExpense(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.DeleteExpense(uint(id)); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "expense deleted", nil)
}
