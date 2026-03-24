package handlers

import (
	"Synconomics/dto"
	"Synconomics/pkg/helpers"
	"Synconomics/services"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/jinzhu/copier"
	"github.com/markbates/goth/gothic"
)

type AuthHandler struct {
	authService     services.AuthService
	businessService services.BusinessService
}

func NewAuthHandler(authService services.AuthService, businessService services.BusinessService) *AuthHandler {
	return &AuthHandler{authService, businessService}
}

// GetMeWithBusinesses
// @Summary Mendapatkan profil user dan daftar bisnisnya
// @Description Mengembalikan data user dan semua bisnis yang dimilikinya berdasarkan token JWT
// @Tags auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=dto.UserBusinessResponse}
// @Router /auth/me-with-businesses [get]
func (h *AuthHandler) GetMeWithBusinesses(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return helpers.ErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	user, _, err := h.authService.GetProfile(userID)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "user not found")
	}

	businesses, err := h.businessService.GetBusinessesByUserId(userID)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusInternalServerError, "failed to fetch businesses")
	}

	var userResp dto.UserResponse
	copier.Copy(&userResp, user)

	var bizResp []dto.BusinessResponse
	copier.Copy(&bizResp, businesses)

	return helpers.SuccessResponse(c, fiber.StatusOK, "user profile and businesses fetched", dto.UserBusinessResponse{
		User:       userResp,
		Businesses: bizResp,
	})
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register
// @Summary Register pengguna baru
// @Description Mendaftarkan user baru menggunakan nama, email, dan password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body handlers.RegisterRequest true "Register Data"
// @Success 201 {object} helpers.Response{data=dto.UserResponse}
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	req := new(RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid body request")
	}

	user, token, err := h.authService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var userResp dto.UserResponse
	copier.Copy(&userResp, user)

	return helpers.SuccessResponse(c, fiber.StatusCreated, "registration success", fiber.Map{
		"token": token,
		"user":  userResp,
	})
}

// Login
// @Summary Login pengguna
// @Description Authentikasi user yang ada menghasilkan JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body handlers.LoginRequest true "Login Data"
// @Success 200 {object} helpers.Response{data=dto.UserResponse}
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid body request")
	}

	user, token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusUnauthorized, "invalid credentials")
	}

	var userResp dto.UserResponse
	copier.Copy(&userResp, user)

	return helpers.SuccessResponse(c, fiber.StatusOK, "login success", fiber.Map{
		"token": token,
		"user":  userResp,
	})
}

// GoogleLogin memulai flow OAuth Google.
// Menggunakan adaptor Fiber→net/http karena gothic dirancang untuk net/http.
// @Summary Memulai login Google
// @Description Redirect ke Google OAuth Login
// @Tags auth
// @Router /auth/google [get]
func (h *AuthHandler) GoogleLogin(c *fiber.Ctx) error {
	// gothic.BeginAuthHandler perlu net/http ResponseWriter & Request
	// adaptor.HTTPHandlerFunc mengkonversi net/http handler ke Fiber handler
	handler := adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set query param "provider" agar gothic tahu provider-nya
		q := r.URL.Query()
		q.Set("provider", "google")
		r.URL.RawQuery = q.Encode()
		gothic.BeginAuthHandler(w, r)
	})
	return handler(c)
}

// Profile
// @Summary Mendapatkan profil user yang sedang login
// @Description Mengembalikan data user berdasarkan token JWT
// @Tags auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helpers.Response{data=dto.UserResponse}
// @Router /auth/profile [get]
func (h *AuthHandler) Profile(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return helpers.ErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	user, token, err := h.authService.GetProfile(userID)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusNotFound, "user not found")
	}

	var userResp dto.UserResponse
	copier.Copy(&userResp, user)

	return helpers.SuccessResponse(c, fiber.StatusOK, "profile fetched", fiber.Map{
		"token": token,
		"user":  userResp,
	})
}

// GoogleCallback menangani callback dari Google setelah user authorize.
// @Summary Callback dari Google OAuth
// @Description Endpoint untuk menangani response dari Google
// @Tags auth
// @Router /auth/google/callback [get]
func (h *AuthHandler) GoogleCallback(c *fiber.Ctx) error {
	var googleUser interface{ GetUserID() string }
	_ = googleUser

	var completeErr error
	var token string

	handler := adaptor.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		q.Set("provider", "google")
		r.URL.RawQuery = q.Encode()

		gu, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			completeErr = err
			return
		}

		_, t, err := h.authService.HandleGoogleCallback(gu)
		if err != nil {
			completeErr = err
			return
		}
		token = t
	})

	if err := handler(c); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if completeErr != nil {
		return c.Status(401).JSON(fiber.Map{"error": "google authentication failed: " + completeErr.Error()})
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	return c.Redirect(frontendURL + "/auth/callback?token=" + token)
}
// UpdateProfile
// @Summary Memperbarui profil user
// @Description Mengubah nama atau email user yang sedang login
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.UpdateProfileRequest true "Update Profile Data"
// @Success 200 {object} helpers.Response{data=dto.UserResponse}
// @Router /auth/profile [put]
func (h *AuthHandler) UpdateProfile(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return helpers.ErrorResponse(c, fiber.StatusUnauthorized, "unauthorized")
	}

	var req dto.UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, "invalid request body")
	}

	user, err := h.authService.UpdateProfile(userID, req.Name, req.Email)
	if err != nil {
		return helpers.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var resp dto.UserResponse
	copier.Copy(&resp, user)

	return helpers.SuccessResponse(c, fiber.StatusOK, "profile updated", resp)
}
