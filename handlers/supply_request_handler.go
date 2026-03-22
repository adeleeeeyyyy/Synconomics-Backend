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

type SupplyRequestHandler struct {
	service services.SupplyRequestService
}

func NewSupplyRequestHandler(service services.SupplyRequestService) *SupplyRequestHandler {
	return &SupplyRequestHandler{service}
}

// CreateSupplyRequest
// @Summary Membuat Supply Request
// @Description Menambahkan data permintaan supply baru
// @Tags supply_requests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateSupplyRequestReq true "Supply Request Body"
// @Success 201 {object} helpers.Response{data=models.SupplyRequest}
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /supply-requests [post]
func (h *SupplyRequestHandler) CreateSupplyRequest(c *fiber.Ctx) error {
	var req dto.CreateSupplyRequestReq
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	var request models.SupplyRequest
	if err := copier.Copy(&request, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.CreateSupplyRequest(&request); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusCreated, "supply request created", request)
}

// GetSupplyRequests
// @Summary Daftar seluruh Supply Request
// @Description Mengambil semua data supply request
// @Tags supply_requests
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]models.SupplyRequest}
// @Failure 500 {object} helpers.Response
// @Router /supply-requests [get]
func (h *SupplyRequestHandler) GetSupplyRequests(c *fiber.Ctx) error {
	requests, err := h.service.GetAllSupplyRequests()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "supply requests fetched", requests)
}

// GetSupplyRequest
// @Summary Detail Supply Request
// @Description Mengambil informasi supply request menggunakan ID
// @Tags supply_requests
// @Produce json
// @Security BearerAuth
// @Param id path int true "Supply Request ID"
// @Success 200 {object} helpers.Response{data=models.SupplyRequest}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /supply-requests/{id} [get]
func (h *SupplyRequestHandler) GetSupplyRequest(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	request, err := h.service.GetSupplyRequestById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "supply request not found")
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "supply request fetched", request)
}

// GetSupplyRequestsByBusinessId
// @Summary Supply Request berdasarkan Business
// @Description Mengambil permintaan supply milik satu bisnis
// @Tags supply_requests
// @Produce json
// @Security BearerAuth
// @Param id path int true "Business ID"
// @Success 200 {object} helpers.Response{data=[]models.SupplyRequest}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /supply-requests/business/{id} [get]
func (h *SupplyRequestHandler) GetSupplyRequestsByBusinessId(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid business id")
	}

	requests, err := h.service.GetSupplyRequestsByBusinessId(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "supply requests not found")
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "supply requests fetched", requests)
}

// UpdateSupplyRequest
// @Summary Memperbarui Supply Request
// @Description Mengubah kolom pada supply request
// @Tags supply_requests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Supply Request ID"
// @Param request body dto.UpdateSupplyRequestReq true "Update Request Body"
// @Success 200 {object} helpers.Response{data=models.SupplyRequest}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /supply-requests/{id} [put]
func (h *SupplyRequestHandler) UpdateSupplyRequest(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	existingReq, err := h.service.GetSupplyRequestById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "supply request not found")
	}

	var req dto.UpdateSupplyRequestReq
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := copier.Copy(existingReq, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.UpdateSupplyRequest(existingReq); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "supply request updated", existingReq)
}

// DeleteSupplyRequest
// @Summary Menghapus Supply Request
// @Description Menghapus permintaan supply
// @Tags supply_requests
// @Produce json
// @Security BearerAuth
// @Param id path int true "Supply Request ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /supply-requests/{id} [delete]
func (h *SupplyRequestHandler) DeleteSupplyRequest(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.DeleteSupplyRequest(uint(id)); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "supply request deleted", nil)
}
