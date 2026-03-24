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

type BusinessScoreHandler struct {
	service services.BusinessScoreService
}

func NewBusinessScoreHandler(service services.BusinessScoreService) *BusinessScoreHandler {
	return &BusinessScoreHandler{service}
}

// CreateScore
// @Summary Create Business Score
// @Description Add a new business score record
// @Tags business_scores
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateBusinessScoreRequest true "Business Score Body"
// @Success 201 {object} helpers.Response{data=dto.BusinessScoreResponse}
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /business-scores [post]
func (h *BusinessScoreHandler) CreateScore(c *fiber.Ctx) error {
	var req dto.CreateBusinessScoreRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	var score models.BusinessScore
	if err := copier.Copy(&score, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.CreateScore(&score); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var res dto.BusinessScoreResponse
	copier.Copy(&res, &score)
	res.ID = score.ID
	res.CreatedAt = score.CreatedAt

	return helpers.SuccessResponse(c, fiber.StatusCreated, "business score created", res)
}

// GetScores
// @Summary List all Business Scores
// @Description Fetch all business score records
// @Tags business_scores
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]dto.BusinessScoreResponse}
// @Failure 500 {object} helpers.Response
// @Router /business-scores [get]
func (h *BusinessScoreHandler) GetScores(c *fiber.Ctx) error {
	scores, err := h.service.GetAllScores()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var res []dto.BusinessScoreResponse
	copier.Copy(&res, &scores)
	for i, s := range scores {
		res[i].ID = s.ID
		res[i].CreatedAt = s.CreatedAt
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "business scores fetched", res)
}

// GetScore
// @Summary Detail Business Score
// @Description Fetch a business score record by ID
// @Tags business_scores
// @Produce json
// @Security BearerAuth
// @Param id path int true "Score ID"
// @Success 200 {object} helpers.Response{data=dto.BusinessScoreResponse}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /business-scores/{id} [get]
func (h *BusinessScoreHandler) GetScore(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	score, err := h.service.GetScoreById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "business score not found")
	}

	var res dto.BusinessScoreResponse
	copier.Copy(&res, score)
	res.ID = score.ID
	res.CreatedAt = score.CreatedAt

	return helpers.SuccessResponse(c, fiber.StatusOK, "business score fetched", res)
}

// GetScoresByBusinessId
// @Summary Business Scores by Business ID
// @Description Fetch business scores for a specific business
// @Tags business_scores
// @Produce json
// @Security BearerAuth
// @Param id path int true "Business ID"
// @Success 200 {object} helpers.Response{data=[]dto.BusinessScoreResponse}
// @Failure 400 {object} helpers.Response
// @Router /business-scores/business/{id} [get]
func (h *BusinessScoreHandler) GetScoresByBusinessId(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid business id")
	}

	scores, err := h.service.GetScoresByBusinessId(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "business scores not found")
	}

	var res []dto.BusinessScoreResponse
	copier.Copy(&res, &scores)
	for i, s := range scores {
		res[i].ID = s.ID
		res[i].CreatedAt = s.CreatedAt
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "business scores fetched", res)
}

// UpdateScore
// @Summary Update Business Score
// @Description Update fields of a business score record
// @Tags business_scores
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Score ID"
// @Param request body dto.UpdateBusinessScoreRequest true "Update Body"
// @Success 200 {object} helpers.Response{data=dto.BusinessScoreResponse}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /business-scores/{id} [put]
func (h *BusinessScoreHandler) UpdateScore(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	existing, err := h.service.GetScoreById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "business score not found")
	}

	var req dto.UpdateBusinessScoreRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := copier.Copy(existing, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.UpdateScore(existing); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var res dto.BusinessScoreResponse
	copier.Copy(&res, existing)
	res.ID = existing.ID
	res.CreatedAt = existing.CreatedAt

	return helpers.SuccessResponse(c, fiber.StatusOK, "business score updated", res)
}

// DeleteScore
// @Summary Delete Business Score
// @Description Remove a business score record
// @Tags business_scores
// @Produce json
// @Security BearerAuth
// @Param id path int true "Score ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Router /business-scores/{id} [delete]
func (h *BusinessScoreHandler) DeleteScore(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.DeleteScore(uint(id)); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "business score deleted", nil)
}

// GetLatestScoreByBusinessId
// @Summary Latest Business Score
// @Description Fetch the most recent business score for a business
// @Tags business_scores
// @Produce json
// @Security BearerAuth
// @Param id path int true "Business ID"
// @Success 200 {object} helpers.Response{data=dto.BusinessScoreResponse}
// @Failure 400 {object} helpers.Response
// @Router /business-scores/business/{id}/latest [get]
func (h *BusinessScoreHandler) GetLatestScoreByBusinessId(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid business id")
	}

	score, err := h.service.GetLatestScoreByBusinessId(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "latest score not found")
	}

	var res dto.BusinessScoreResponse
	copier.Copy(&res, score)
	res.ID = score.ID
	res.CreatedAt = score.CreatedAt

	return helpers.SuccessResponse(c, fiber.StatusOK, "latest business score fetched", res)
}
