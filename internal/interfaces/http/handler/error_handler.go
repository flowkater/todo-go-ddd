package handler

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ErrorHandler struct {
	logger *zap.Logger
}

func NewErrorHandler(logger *zap.Logger) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
	}
}

func (h *ErrorHandler) HandleError(c *fiber.Ctx, err error, status int, msg string, fields ...zap.Field) error {
	if msg == "" {
		msg = err.Error()
	}
	
	// 에러 로깅
	if status >= fiber.StatusInternalServerError {
		h.logger.Error(msg, append(fields, zap.Error(err))...)
	} else {
		h.logger.Warn(msg, append(fields, zap.Error(err))...)
	}

	return c.Status(status).JSON(fiber.Map{
		"error": msg,
	})
}
