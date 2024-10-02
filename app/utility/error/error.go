package error_utils

import "github.com/gofiber/fiber/v2"

var (
	ErrorNotFound       = fiber.NewError(fiber.StatusNotFound, "Not found")
	ErrorBadRequest     = fiber.NewError(fiber.StatusBadRequest, "Bad request")
	ErrorUnauthorized   = fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	ErrorConflict       = fiber.NewError(fiber.StatusConflict, "Conflict")
	ErrorInternalServer = fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
)
