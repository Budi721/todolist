package web

import (
	"github.com/gofiber/fiber/v2"
)

func InternalServerErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal server error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	return ctx.Status(code).JSON(Response{
		Status:  message,
		Message: message,
		Data:    nil,
	})
}
