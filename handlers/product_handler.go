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

type ProductHandler struct {
	service         services.ProductService
	businessService services.BusinessService
}

func NewProductHandler(service services.ProductService, businessService services.BusinessService) *ProductHandler {
	return &ProductHandler{service, businessService}
}

// CreateProduct
// @Summary Membuat produk baru
// @Description Menambahkan data entitas produk baru dan menyimpan path gambar
// @Tags products
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param business_id formData integer true "Business ID"
// @Param name formData string true "Product Name"
// @Param description formData string false "Product Description"
// @Param price formData number true "Price"
// @Param stock formData integer false "Stock"
// @Param image_url formData file true "Product Image"
// @Success 201 {object} helpers.Response{data=dto.ProductResponse}
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var req dto.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	_, filePath, err := helpers.HandleFileUpload(c, "image_url", "./public/uploads")
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var newProduct models.Product
	if err := copier.Copy(&newProduct, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to map")
	}

	// Validate Business existence
	_, err = h.businessService.GetBusinessById(newProduct.BusinessID)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "business not found: please ensure the business_id is correct")
	}

	newProduct.ImageURL = filePath

	if err := h.service.CreateProduct(&newProduct); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp dto.ProductResponse
	copier.Copy(&resp, &newProduct)

	return helpers.SuccessResponse(c, fiber.StatusCreated, "product created", resp)
}

// GetProducts
// @Summary ambil semua daftar produk
// @Description mengambil semua data produk yang ada
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} helpers.Response{data=[]dto.ProductResponse} "Sukses mengambil data"
// @Failure 500 {object} helpers.Response "internal server error"
// @Router /products [get]
func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.service.GetAllProducts()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp []dto.ProductResponse
	copier.Copy(&resp, &products)

	return helpers.SuccessResponse(c, fiber.StatusOK, "products fetched", resp)
}

// GetProduct
// @Summary Mendapatkan detail produk
// @Description Mengambil informasi lengkap terkait satu produk menggunakan ID
// @Tags products
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} helpers.Response{data=dto.ProductResponse}
// @Router /products/{id} [get]
func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	product, err := h.service.GetProductById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "product not found")
	}

	var resp dto.ProductResponse
	copier.Copy(&resp, product)

	return helpers.SuccessResponse(c, fiber.StatusOK, "product fetched", resp)
}

// GetProductsByBusiness
// @Summary Mendapatkan daftar produk dalam satu bisnis
// @Description Mengambil semua data produk milik sebuah bisnis
// @Tags products
// @Produce json
// @Security BearerAuth
// @Param businessId path int true "Business ID"
// @Success 200 {object} helpers.Response{data=[]dto.ProductResponse}
// @Router /products/business/{businessId} [get]
func (h *ProductHandler) GetProductsByBusiness(c *fiber.Ctx) error {
	businessId, err := strconv.ParseUint(c.Params("businessId"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid business id")
	}

	products, err := h.service.GetProductsByBusinessId(uint(businessId))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp []dto.ProductResponse
	copier.Copy(&resp, &products)

	return helpers.SuccessResponse(c, fiber.StatusOK, "business products fetched", resp)
}

// UpdateProduct
// @Summary Memperbarui parameter produk
// @Description Mengubah kolom dari sebuah entitas produk
// @Tags products
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param name formData string false "Product Name"
// @Param description formData string false "Product Description"
// @Param price formData number false "Price"
// @Param stock formData integer false "Stock"
// @Success 200 {object} helpers.Response{data=dto.ProductResponse}
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	existingProduct, err := h.service.GetProductById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "product not found")
	}

	if err := c.BodyParser(&existingProduct); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid form input")
	}

	if err := h.service.UpdateProduct(existingProduct); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp dto.ProductResponse
	copier.Copy(&resp, existingProduct)

	return helpers.SuccessResponse(c, fiber.StatusOK, "product updated", resp)
}

// DeleteProduct
// @Summary Menghapus produk
// @Description Menghapus produk dari database
// @Tags products
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Product ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.DeleteProduct(uint(id)); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "product deleted", nil)
}
