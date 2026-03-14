package handlers

import (
    "github.com/gofiber/fiber/v2"
    "Synconomics/models"
    "Synconomics/pkg"
    "Synconomics/services"
    "strings"
)

type UserHandler struct {
    service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
    var req models.RegisterRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "request tidak valid",
        })
    }

    user, err := h.service.Register(req)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "registrasi berhasil",
        "data":    user,
    })
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
    var req models.LoginRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "request tidak valid",
        })
    }

    token, err := h.service.Login(req)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "message": "login berhasil",
        "token":   token,
    })
}

func (h *UserHandler) Profile(c *fiber.Ctx) error {
    userID := c.Locals("userID").(uint)

    user, err := h.service.GetProfile(userID)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "user tidak ditemukan",
        })
    }

    return c.JSON(fiber.Map{
        "data": user,
    })
}

// JWTMiddleware — cek token di setiap protected route
func JWTMiddleware(c *fiber.Ctx) error {
    auth := c.Get("Authorization")
    if !strings.HasPrefix(auth, "Bearer ") {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "token tidak ada",
        })
    }

    claims, err := pkg.ValidateToken(strings.TrimPrefix(auth, "Bearer "))
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "token tidak valid atau expired",
        })
    }

    c.Locals("userID", claims.UserID)
    return c.Next()
}