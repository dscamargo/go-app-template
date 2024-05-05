package web

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ErrorResponse struct {
	Messages []string `json:"message"`
}

func NewInternalServerError(c *fiber.Ctx) error {
	return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": "Internal Server Error"})
}

func NewOKResponse(c *fiber.Ctx, data any) error {
	return c.Status(http.StatusOK).JSON(data)
}

func NewNoContentResponse(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNoContent)
}

func NewUnauthorizedResponse(c *fiber.Ctx, messages []string) error {
	return newGenericErrorResponse(c, 401, messages)
}

func newGenericErrorResponse(c *fiber.Ctx, code int, messages []string) error {
	return c.Status(code).JSON(ErrorResponse{Messages: messages})
}
