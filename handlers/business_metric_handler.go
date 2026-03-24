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

type BusinessMetricHandler struct {
	service services.BusinessMetricService
}

func NewBusinessMetricHandler(service services.BusinessMetricService) *BusinessMetricHandler {
	return &BusinessMetricHandler{service}
}

// CreateMetric
// @Summary Create Business Metric
// @Description Add a new business metric record
// @Tags business_metrics
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateBusinessMetricRequest true "Business Metric Body"
// @Success 201 {object} helpers.Response{data=dto.BusinessMetricResponse}
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /business-metrics [post]
func (h *BusinessMetricHandler) CreateMetric(c *fiber.Ctx) error {
	var req dto.CreateBusinessMetricRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	var metric models.BusinessMetric
	if err := copier.Copy(&metric, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.CreateMetric(&metric); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var res dto.BusinessMetricResponse
	copier.Copy(&res, &metric)
	res.ID = metric.ID
	res.CreatedAt = metric.CreatedAt

	return helpers.SuccessResponse(c, fiber.StatusCreated, "business metric created", res)
}

// GetMetrics
// @Summary List all Business Metrics
// @Description Fetch all business metric records
// @Tags business_metrics
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]dto.BusinessMetricResponse}
// @Failure 500 {object} helpers.Response
// @Router /business-metrics [get]
func (h *BusinessMetricHandler) GetMetrics(c *fiber.Ctx) error {
	metrics, err := h.service.GetAllMetrics()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var res []dto.BusinessMetricResponse
	copier.Copy(&res, &metrics)
	for i, m := range metrics {
		res[i].ID = m.ID
		res[i].CreatedAt = m.CreatedAt
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "business metrics fetched", res)
}

// GetMetric
// @Summary Detail Business Metric
// @Description Fetch a business metric record by ID
// @Tags business_metrics
// @Produce json
// @Security BearerAuth
// @Param id path int true "Metric ID"
// @Success 200 {object} helpers.Response{data=dto.BusinessMetricResponse}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /business-metrics/{id} [get]
func (h *BusinessMetricHandler) GetMetric(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	metric, err := h.service.GetMetricById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "business metric not found")
	}

	var res dto.BusinessMetricResponse
	copier.Copy(&res, metric)
	res.ID = metric.ID
	res.CreatedAt = metric.CreatedAt

	return helpers.SuccessResponse(c, fiber.StatusOK, "business metric fetched", res)
}

// GetMetricsByBusinessId
// @Summary Business Metrics by Business ID
// @Description Fetch business metrics for a specific business
// @Tags business_metrics
// @Produce json
// @Security BearerAuth
// @Param id path int true "Business ID"
// @Success 200 {object} helpers.Response{data=[]dto.BusinessMetricResponse}
// @Failure 400 {object} helpers.Response
// @Router /business-metrics/business/{id} [get]
func (h *BusinessMetricHandler) GetMetricsByBusinessId(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid business id")
	}

	metrics, err := h.service.GetMetricsByBusinessId(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "business metrics not found")
	}

	var res []dto.BusinessMetricResponse
	copier.Copy(&res, &metrics)
	for i, m := range metrics {
		res[i].ID = m.ID
		res[i].CreatedAt = m.CreatedAt
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "business metrics fetched", res)
}

// UpdateMetric
// @Summary Update Business Metric
// @Description Update fields of a business metric record
// @Tags business_metrics
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Metric ID"
// @Param request body dto.UpdateBusinessMetricRequest true "Update Body"
// @Success 200 {object} helpers.Response{data=dto.BusinessMetricResponse}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /business-metrics/{id} [put]
func (h *BusinessMetricHandler) UpdateMetric(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	existing, err := h.service.GetMetricById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "business metric not found")
	}

	var req dto.UpdateBusinessMetricRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := copier.Copy(existing, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.UpdateMetric(existing); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var res dto.BusinessMetricResponse
	copier.Copy(&res, existing)
	res.ID = existing.ID
	res.CreatedAt = existing.CreatedAt

	return helpers.SuccessResponse(c, fiber.StatusOK, "business metric updated", res)
}

// DeleteMetric
// @Summary Delete Business Metric
// @Description Remove a business metric record
// @Tags business_metrics
// @Produce json
// @Security BearerAuth
// @Param id path int true "Metric ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Router /business-metrics/{id} [delete]
func (h *BusinessMetricHandler) DeleteMetric(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.DeleteMetric(uint(id)); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "business metric deleted", nil)
}

// GetLatestMetricByBusinessId
// @Summary Latest Business Metric
// @Description Fetch the most recent business metric for a business
// @Tags business_metrics
// @Produce json
// @Security BearerAuth
// @Param id path int true "Business ID"
// @Success 200 {object} helpers.Response{data=dto.BusinessMetricResponse}
// @Failure 400 {object} helpers.Response
// @Router /business-metrics/business/{id}/latest [get]
func (h *BusinessMetricHandler) GetLatestMetricByBusinessId(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid business id")
	}

	metric, err := h.service.GetLatestMetricByBusinessId(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "latest metric not found")
	}

	var res dto.BusinessMetricResponse
	copier.Copy(&res, metric)
	res.ID = metric.ID
	res.CreatedAt = metric.CreatedAt

	return helpers.SuccessResponse(c, fiber.StatusOK, "latest business metric fetched", res)
}
