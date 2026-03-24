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

type SupplyMatchHandler struct {
	service services.SupplyMatchService
}

func NewSupplyMatchHandler(service services.SupplyMatchService) *SupplyMatchHandler {
	return &SupplyMatchHandler{service}
}

// CreateSupplyMatch
// @Summary Membuat Supply Match
// @Description Menghubungkan supply request dengan supply offer
// @Tags supply_matches
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateSupplyMatchReq true "Supply Match Body"
// @Success 201 {object} helpers.Response{data=models.SupplyMatch}
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /supply-matches [post]
func (h *SupplyMatchHandler) CreateSupplyMatch(c *fiber.Ctx) error {
	var req dto.CreateSupplyMatchReq
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	var match models.SupplyMatch
	if err := copier.Copy(&match, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.CreateSupplyMatch(&match); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp dto.SupplyMatchResponse
	copier.Copy(&resp, &match)

	return helpers.SuccessResponse(c, fiber.StatusCreated, "supply match created", resp)
}

// GetSupplyMatches
// @Summary Daftar seluruh Supply Match
// @Description Mengambil semua data supply match
// @Tags supply_matches
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]dto.SupplyMatchResponse}
// @Failure 500 {object} helpers.Response
// @Router /supply-matches [get]
func (h *SupplyMatchHandler) GetSupplyMatches(c *fiber.Ctx) error {
	matches, err := h.service.GetAllSupplyMatches()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp []dto.SupplyMatchResponse
	copier.Copy(&resp, &matches)

	return helpers.SuccessResponse(c, fiber.StatusOK, "supply matches fetched", resp)
}

// GetSupplyMatch
// @Summary Detail Supply Match
// @Description Mengambil informasi supply match menggunakan ID
// @Tags supply_matches
// @Produce json
// @Security BearerAuth
// @Param id path int true "Supply Match ID"
// @Success 200 {object} helpers.Response{data=dto.SupplyMatchResponse}
// @Router /supply-matches/{id} [get]
func (h *SupplyMatchHandler) GetSupplyMatch(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	match, err := h.service.GetSupplyMatchById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "supply match not found")
	}

	var resp dto.SupplyMatchResponse
	copier.Copy(&resp, match)

	return helpers.SuccessResponse(c, fiber.StatusOK, "supply match fetched", resp)
}

// UpdateSupplyMatchStatus
// @Summary Update status Supply Match
// @Description Mengubah status (accepted/rejected) supply match
// @Tags supply_matches
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Supply Match ID"
// @Param request body dto.UpdateSupplyMatchStatusReq true "Update Status Body"
// @Success 200 {object} helpers.Response{data=dto.SupplyMatchResponse}
// @Router /supply-matches/{id}/status [patch]
func (h *SupplyMatchHandler) UpdateSupplyMatchStatus(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	match, err := h.service.GetSupplyMatchById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "supply match not found")
	}

	var req dto.UpdateSupplyMatchStatusReq
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	match.Status = models.MatchStatus(req.Status)

	if err := h.service.UpdateSupplyMatch(match); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp dto.SupplyMatchResponse
	copier.Copy(&resp, match)

	return helpers.SuccessResponse(c, fiber.StatusOK, "supply match status updated", resp)
}

// DeleteSupplyMatch
// @Summary Menghapus Supply Match
// @Description Menghapus data supply match
// @Tags supply_matches
// @Produce json
// @Security BearerAuth
// @Param id path int true "Supply Match ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /supply-matches/{id} [delete]
func (h *SupplyMatchHandler) DeleteSupplyMatch(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.DeleteSupplyMatch(uint(id)); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "supply match deleted", nil)
}
