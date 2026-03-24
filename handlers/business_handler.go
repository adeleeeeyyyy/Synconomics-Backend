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

type BusinessHandler struct {
	businessService services.BusinessService
}

func NewBusinessHandler(bService services.BusinessService) *BusinessHandler {
	return &BusinessHandler{
		businessService: bService,
	}
}

// CreateBusiness handles business creation with Logo upload and Copier mapping
// @Summary Membuat profil bisnis baru
// @Description Membuat entitas bisnis baru dengan data dan logo
// @Tags business
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param name formData string true "Business Name"
// @Param description formData string false "Business Description"
// @Param category formData string true "Business Category"
// @Param address formData string false "Business Address"
// @Param latitude formData number false "Latitude"
// @Param longitude formData number false "Longitude"
// @Param phone formData string false "Phone Number"
// @Param whatsapp formData string false "WhatsApp"
// @Param instagram formData string false "Instagram"
// @Param tiktok formData string false "Tiktok"
// @Param website formData string false "Website"
// @Param logo_url formData file true "Business Logo"
// @Success 201 {object} helpers.Response{data=models.Business}
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /business [post]
func (h *BusinessHandler) CreateBusiness(c *fiber.Ctx) error {
	// 1. Ambil data DTO dari request form
	var req dto.CreateBusinessRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "Format form-data tidak valid")
	}

	// 2. Upload Logo Image (Optional/Wajib tergantung kebutuhan, di sini dibuat wajib)
	_, logoPath, err := helpers.HandleFileUpload(c, "logo_url", "./public/uploads/logos")
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "Logo wajib disertakan dan berupa gambar: "+err.Error())
	}

	// 3. Persiapkan variabel Model Database
	var business models.Business

	// 4. Salin / Mapping nilai DTO persis ke Model `models.Business` menggunakan Copier
	if err := copier.Copy(&business, &req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memproses/mapping data DTO")
	}

	// 5. Masukkan path logo yang baru selesai diupload
	business.LogoURL = logoPath

	// 6. Masukkan UserID (Dari JWT middleware lokal)
	userID := c.Locals("userID")
	if userID != nil {
		if id, ok := userID.(uint); ok {
			business.UserID = id
		} else if idFloat, ok := userID.(float64); ok {
			business.UserID = uint(idFloat)
		}
	}

	// 7. Panggil layer service dengan entitas Model penuh
	err = h.businessService.CreateBusiness(&business)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	var resp dto.BusinessResponse
	copier.Copy(&resp, &business)

	return helpers.SuccessResponse(c, fiber.StatusCreated, "Business berhasil dibuat", resp)
}

// GetAllBusinesses
// @Summary Mengambil semua data bisnis
// @Description Menampilkan daftar seluruh bisnis yang terdaftar
// @Tags business
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]dto.BusinessResponse}
// @Failure 500 {object} helpers.Response
// @Router /business [get]
func (h *BusinessHandler) GetAllBusinesses(c *fiber.Ctx) error {
	businesses, err := h.businessService.GetAllBusinesses()
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data bisnis")
	}

	var resp []dto.BusinessResponse
	copier.Copy(&resp, &businesses)

	return helpers.SuccessResponse(c, fiber.StatusOK, "Berhasil mengambil semua data bisnis", resp)
}

// GetBusinessById
// @Summary Mengambil data bisnis dengan ID
// @Description Mendapatkan detil sebuah bisnis menggunakan path ID
// @Tags business
// @Produce json
// @Security BearerAuth
// @Param id path int true "Business ID"
// @Success 200 {object} helpers.Response{data=dto.BusinessResponse}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Router /business/{id} [get]
func (h *BusinessHandler) GetBusinessById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	business, err := h.businessService.GetBusinessById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "Data bisnis tidak ditemukan")
	}

	var resp dto.BusinessResponse
	copier.Copy(&resp, business)

	return helpers.SuccessResponse(c, fiber.StatusOK, "Berhasil mengambil data bisnis", resp)
}

// GetMyBusinesses
// @Summary Mengambil bisnis milik user login
// @Description Mendapatkan daftar bisnis yang dimiliki oleh user yang sedang login (berdasarkan ID dari JWT)
// @Tags business
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=[]dto.BusinessResponse}
// @Failure 500 {object} helpers.Response
// @Router /business/me [get]
func (h *BusinessHandler) GetMyBusinesses(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	businesses, err := h.businessService.GetBusinessesByUserId(userID)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data bisnis user")
	}

	var resp []dto.BusinessResponse
	copier.Copy(&resp, &businesses)

	return helpers.SuccessResponse(c, fiber.StatusOK, "Berhasil mengambil bisnis Anda", resp)
}

// UpdateBusiness handles updating existing business
// @Summary Memperbarui bisnis yang sudah ada
// @Description Mengubah sebagian atau seluruh data bisnis menggunakan ID beserta opsional pembaruan logo
// @Tags business
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param id path int true "Business ID"
// @Param name formData string false "Business Name"
// @Param description formData string false "Business Description"
// @Param category formData string false "Business Category"
// @Param address formData string false "Business Address"
// @Param latitude formData number false "Latitude"
// @Param longitude formData number false "Longitude"
// @Param phone formData string false "Phone Number"
// @Param whatsapp formData string false "WhatsApp"
// @Param instagram formData string false "Instagram"
// @Param tiktok formData string false "Tiktok"
// @Param website formData string false "Website"
// @Param logo_url formData file false "Business Logo (Optional)"
// @Success 200 {object} helpers.Response{data=dto.BusinessResponse}
// @Failure 400 {object} helpers.Response
// @Failure 404 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /business/{id} [put]
func (h *BusinessHandler) UpdateBusiness(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "ID bisnis tidak valid")
	}

	// 1. Cek existensi business
	existingBusiness, err := h.businessService.GetBusinessById(uint(id))
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "Data bisnis tidak ditemukan")
	}

	// 2. Parse DTO Update
	var req dto.UpdateBusinessRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "Format form-data tidak valid")
	}

	// 3. Upload Logo jika ada / diperbarui (TIDAK WAJIB)
	_, logoPath, errUpload := helpers.HandleFileUpload(c, "logo_url", "./public/uploads/logos")
	if errUpload == nil && logoPath != "" {
		existingBusiness.LogoURL = logoPath // Timpa hanya jika ada form part gambar
	}

	// 4. Copier: timpa field dari req (DTO) ke existingBusiness.
	// Gunakan IgnoreEmpty agar field kosongan dari request tidak menimpa data aslinya di Database
	if err := copier.CopyWithOption(existingBusiness, &req, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memproses data update")
	}

	// 5. Update di Service
	if err := h.businessService.UpdateBusiness(existingBusiness); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memperbarui data: "+err.Error())
	}

	var resp dto.BusinessResponse
	copier.Copy(&resp, existingBusiness)

	return helpers.SuccessResponse(c, fiber.StatusOK, "Data bisnis berhasil diperbarui", resp)
}

// DeleteBusiness handles deleting a business
// @Summary Menghapus bisnis berdasarkan ID
// @Description Menghapus record dari ID entitas bisnis yang ditentukan
// @Tags business
// @Produce json
// @Security BearerAuth
// @Param id path int true "Business ID"
// @Success 200 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Router /business/{id} [delete]
func (h *BusinessHandler) DeleteBusiness(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "ID bisnis tidak valid")
	}

	// Opsional: Cek existensi via Service atau langsung eksekusi Delete
	if err := h.businessService.DeleteBusiness(uint(id)); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus data bisnis: "+err.Error())
	}

	return helpers.SuccessResponse(c, fiber.StatusOK, "Bisnis berhasil dihapus", nil)
}
