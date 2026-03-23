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

type SupplyOfferHandler struct {
	service services.SupplyOfferService
}

func NewSupplyOfferHandler(service services.SupplyOfferService) *SupplyOfferHandler {
	return &SupplyOfferHandler{service}
}

// CreateSupplyOffer
// @Summary Membuat Supply Offer
// @Description Menambahkan data penawaran supply baru
// @Tags supply_offers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateSupplyOfferReq true "Supply Offer Body"
// @Success 201 {object} helpers.Response{data=models.SupplyOffer}
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /supply-offers [post]
func (h *SupplyOfferHandler) CreateSupplyOffer(c *fiber.Ctx) error {
	var req dto.CreateSupplyOfferReq
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	var offer models.SupplyOffer
	if err := copier.Copy(&offer, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.CreateSupplyOffer(&offer); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusCreated, "supply offer created", offer)
}

// GetSupplyOffers
// @Summary Daftar seluruh Supply Offer
// @Description Mengambil semua data supply offer
// @Tags supply_offers
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]models.SupplyOffer}
// @Failure 500 {object} helpers.Response
// @Router /supply-offers [get]
func (h *SupplyOfferHandler) GetSupplyOffers(c *fiber.Ctx) error {
	offers, err := h.service.GetAllSupplyOffers()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "supply offers fetched", offers)
}

// GetSupplyOffer
// @Summary Detail Supply Offer
// @Description Mengambil informasi supply offer menggunakan ID
// @Tags supply_offers
// @Produce json
// @Security BearerAuth
// @Param id path int true "Supply Offer ID"
// @Success 200 {object} helpers.Response{data=models.SupplyOffer}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /supply-offers/{id} [get]
func (h *SupplyOfferHandler) GetSupplyOffer(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	offer, err := h.service.GetSupplyOfferById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "supply offer not found")
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "supply offer fetched", offer)
}

// GetSupplyOffersByBusinessId
// @Summary Supply Offer berdasarkan Business
// @Description Mengambil penawaran supply milik satu bisnis
// @Tags supply_offers
// @Produce json
// @Security BearerAuth
// @Param id path int true "Business ID"
// @Success 200 {object} helpers.Response{data=[]models.SupplyOffer}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /supply-offers/business/{id} [get]
func (h *SupplyOfferHandler) GetSupplyOffersByBusinessId(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid business id")
	}

	offers, err := h.service.GetSupplyOffersByBusinessId(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "supply offers not found")
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "supply offers fetched", offers)
}

// UpdateSupplyOffer
// @Summary Memperbarui Supply Offer
// @Description Mengubah kolom pada supply offer
// @Tags supply_offers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Supply Offer ID"
// @Param request body dto.UpdateSupplyOfferReq true "Update Request Body"
// @Success 200 {object} helpers.Response{data=models.SupplyOffer}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /supply-offers/{id} [put]
func (h *SupplyOfferHandler) UpdateSupplyOffer(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	existingOffer, err := h.service.GetSupplyOfferById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "supply offer not found")
	}

	var req dto.UpdateSupplyOfferReq
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := copier.Copy(existingOffer, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	if err := h.service.UpdateSupplyOffer(existingOffer); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "supply offer updated", existingOffer)
}

// DeleteSupplyOffer
// @Summary Menghapus Supply Offer
// @Description Menghapus penawaran supply
// @Tags supply_offers
// @Produce json
// @Security BearerAuth
// @Param id path int true "Supply Offer ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /supply-offers/{id} [delete]
func (h *SupplyOfferHandler) DeleteSupplyOffer(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.DeleteSupplyOffer(uint(id)); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "supply offer deleted", nil)
}
