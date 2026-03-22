package helpers

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(Response{
		Status: true,
		Message: message,
		Data: data,
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(Response{
		Status: false,
		Message: message,
	})
}
