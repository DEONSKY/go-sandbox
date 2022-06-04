package utils

import (
	"github.com/DEONSKY/go-sandbox/helper"
	"github.com/gofiber/fiber/v2"
)

type httpError struct {
	Statuscode int    `json:"statusCode"`
	Error      string `json:"error"`
}

// ErrorHandler is used to catch error thrown inside the routes by ctx.Next(err)
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError

	// Check if it's an fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})

	return c.Status(code).JSON(res)
}
