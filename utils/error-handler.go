package utils

import (
	"fmt"

	"github.com/DEONSKY/go-sandbox/helper"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func ReturnErrorResponse(code int, message string, errors []string) ErrorResponse {
	return ErrorResponse{Code: code, Message: message, Errors: errors}
}

func ReturnErrorMessage(code int, message string) ErrorResponse {
	return ErrorResponse{Code: code, Message: message, Errors: []string{}}
}

func (c ErrorResponse) Error() string {
	return fmt.Sprintf("Code: %d Message:%s Errors:%v", c.Code, c.Message, c.Errors)
}

// ErrorHandler is used to catch error thrown inside the routes by ctx.Next(err)
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError
	message := "Failed to process request"
	var errors []string

	// Check if it's an fiber.Error type
	if e, ok := err.(ErrorResponse); ok {
		code = e.Code
		message = e.Message
		errors = e.Errors
	}

	res := helper.BuildErrorResponse(message, errors, helper.EmptyObj{})

	return c.Status(code).JSON(res)
}
