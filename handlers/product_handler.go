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
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

// CreateProduct
// @Summary Membuat produk baru
// @Description Menambahkan data entitas produk baru dan menyimpan path gambar
// @Tags products
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param name formData string true "Product Name"
// @Param description formData string false "Product Description"
// @Param price formData number true "Price"
// @Param stock formData integer false "Stock"
// @Param image_url formData file true "Product Image"
// @Success 201 {object} helpers.Response{data=models.Product}
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

	newProduct.ImageURL = filePath

	if err := h.service.CreateProduct(&newProduct); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusCreated, "product creaed", newProduct)
}

// GetProducts
// @Summary ambil semua daftar produk
// @Description mengambil semua data produk yang ada
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} helpers.Response{data=[]models.Product} "Sukses mengambil data"
// @Failure 500 {object} helpers.Response "internal server error"
// @Router /products [get]
func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.service.GetAllProducts()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "products fetched", products)
}

// GetProduct
// @Summary Mendapatkan detail produk
// @Description Mengambil informasi lengkap terkait satu produk menggunakan ID
// @Tags products
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} helpers.Response{data=models.Product}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
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

	return helpers.SuccessResponse(c, fiber.StatusOK, "product fetched", product)
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
// @Success 200 {object} helpers.Response{data=models.Product}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Failure 500 {object} helpers.Response
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

	return helpers.SuccessResponse(c, fiber.StatusOK, "product updated", existingProduct)
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
