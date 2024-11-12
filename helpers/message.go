package helpers

import (
	"github.com/gofiber/fiber/v2"
)

type ResultMessage struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseMessage(c *fiber.Ctx, code int, message string, data interface{}) error {
	response := &ResultMessage{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return c.Status(code).JSON(response)
}
