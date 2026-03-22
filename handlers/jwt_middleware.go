package handlers

import (
	"Synconomics/pkg"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "missing authorization header"})
	}

	// Tangani kasus token dengan auto "Bearer " strip jika ada
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := pkg.ValidateToken(tokenStr)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid or expired token"})
	}

	c.Locals("userID", claims.UserID)
	return c.Next()
}
