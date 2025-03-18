package responses

import (
	"github.com/gofiber/fiber/v2"
	"go-nonce/errs"
	"net/http"
)

var (
	code    int
	message string
)

type ErrorResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

func NewErrorResponses(ctx *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case errs.AppError:
		code = e.Status
		message = e.Message
	case error:
		code = http.StatusUnprocessableEntity
		message = err.Error()
	}
	errorResponse := ErrorResponse{
		Status: false,
		Error:  message,
	}
	return ctx.Status(code).JSON(errorResponse)
}

func NewCustomErrorResponses(ctx *fiber.Ctx, code int, message string) error {
	errorResponse := ErrorResponse{
		Status: false,
		Error:  message,
	}
	return ctx.Status(code).JSON(errorResponse)
}

func NewSuccessResponse(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    data,
	})
}

func NewCreateSuccessResponse(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"data":    data,
		"message": "Data created successfully",
	})
}

func NewErrorBadRequest(ctx *fiber.Ctx, details interface{}) error {
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
		"status": false,
		"error":  details,
	})
}

func NewErrorNotFound(ctx *fiber.Ctx, details interface{}) error {
	return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
		"status": false,
		"error":  details,
	})
}

func NewSuccessMessage(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": data,
	})
}

func NewErrorValidate(ctx *fiber.Ctx, data interface{}) error {
	validateError := fiber.Map{
		"error":  data,
		"status": false,
	}
	return ctx.Status(http.StatusUnprocessableEntity).JSON(validateError)
}
