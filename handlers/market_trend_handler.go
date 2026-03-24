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

type MarketTrendHandler struct {
	service services.MarketTrendService
}

func NewMarketTrendHandler(service services.MarketTrendService) *MarketTrendHandler {
	return &MarketTrendHandler{service}
}

// CreateTrend
// @Summary Membuat data tren pasar baru
// @Description Menambahkan entitas market trend baru
// @Tags market-trends
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param trend body dto.CreateMarketTrendRequest true "Market Trend Data"
// @Success 201 {object} helpers.Response{data=dto.MarketTrendResponse}
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /market-trends [post]
func (h *MarketTrendHandler) CreateTrend(c *fiber.Ctx) error {
	var req dto.CreateMarketTrendRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	var trend models.MarketTrend
	copier.Copy(&trend, &req)

	if err := h.service.CreateTrend(&trend); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp dto.MarketTrendResponse
	copier.Copy(&resp, &trend)

	return helpers.SuccessResponse(c, fiber.StatusCreated, "market trend created", resp)
}

// GetAllTrends
// @Summary Ambil semua tren pasar
// @Description Mengambil semua data tren pasar
// @Tags market-trends
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]dto.MarketTrendResponse}
// @Router /market-trends [get]
func (h *MarketTrendHandler) GetAllTrends(c *fiber.Ctx) error {
	trends, err := h.service.GetAllTrends()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp []dto.MarketTrendResponse
	copier.Copy(&resp, &trends)

	return helpers.SuccessResponse(c, fiber.StatusOK, "market trends fetched", resp)
}

// GetTrendById
// @Summary Ambil tren pasar berdasarkan ID
// @Description Mengambil satu data tren pasar menggunakan ID
// @Tags market-trends
// @Produce json
// @Security BearerAuth
// @Param id path int true "Trend ID"
// @Success 200 {object} helpers.Response{data=dto.MarketTrendResponse}
// @Router /market-trends/{id} [get]
func (h *MarketTrendHandler) GetTrendById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	trend, err := h.service.GetTrendById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "market trend not found")
	}

	var resp dto.MarketTrendResponse
	copier.Copy(&resp, trend)

	return helpers.SuccessResponse(c, fiber.StatusOK, "market trend fetched", resp)
}

// UpdateTrend
// @Summary Perbarui tren pasar
// @Description Mengubah data tren pasar yang sudah ada
// @Tags market-trends
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Trend ID"
// @Param trend body dto.UpdateMarketTrendRequest true "Market Trend Update Data"
// @Success 200 {object} helpers.Response{data=dto.MarketTrendResponse}
// @Router /market-trends/{id} [put]
func (h *MarketTrendHandler) UpdateTrend(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	existingTrend, err := h.service.GetTrendById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "market trend not found")
	}

	var req dto.UpdateMarketTrendRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	copier.CopyWithOption(existingTrend, &req, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	if err := h.service.UpdateTrend(existingTrend); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp dto.MarketTrendResponse
	copier.Copy(&resp, existingTrend)

	return helpers.SuccessResponse(c, fiber.StatusOK, "market trend updated", resp)
}

// DeleteTrend
// @Summary Hapus tren pasar
// @Description Menghapus data tren pasar
// @Tags market-trends
// @Produce json
// @Security BearerAuth
// @Param id path int true "Trend ID"
// @Success 200 {object} helpers.Response
// @Router /market-trends/{id} [delete]
func (h *MarketTrendHandler) DeleteTrend(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.DeleteTrend(uint(id)); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "market trend deleted", nil)
}

// GetTopTrends
// @Summary Ambil 10 tren pasar teratas
// @Description Mengambil 10 data tren pasar dengan skor permintaan tertinggi
// @Tags market-trends
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]dto.MarketTrendResponse}
// @Router /market-trends/top [get]
func (h *MarketTrendHandler) GetTopTrends(c *fiber.Ctx) error {
	trends, err := h.service.GetTopTenTrends()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp []dto.MarketTrendResponse
	copier.Copy(&resp, &trends)

	return helpers.SuccessResponse(c, fiber.StatusOK, "top 10 market trends fetched", resp)
}
