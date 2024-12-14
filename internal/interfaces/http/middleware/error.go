package middleware

import (
	"github.com/flowkater/ddd-todo-app/internal/interfaces/http/errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// ErrorMiddleware creates a new error handling middleware
func ErrorMiddleware(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 다음 핸들러 실행
		err := c.Next()
		if err == nil {
			return nil
		}

		// 에러 타입 확인
		var statusCode int
		var message string
		var internal error

		switch e := err.(type) {
		case *errors.HTTPError:
			statusCode = e.Status()
			message = e.Message
			internal = e.Internal
		case *fiber.Error:
			statusCode = e.Code
			message = e.Message
		default:
			statusCode = fiber.StatusInternalServerError
			message = "Internal Server Error"
			internal = err
		}

		// 에러 로깅
		if statusCode >= fiber.StatusInternalServerError {
			if internal != nil {
				logger.Error(message,
					zap.Error(internal),
					zap.Int("status", statusCode),
				)
			} else {
				logger.Error(message,
					zap.Error(err),
					zap.Int("status", statusCode),
				)
			}
		} else {
			logger.Warn(message,
				zap.Error(err),
				zap.Int("status", statusCode),
			)
		}

		// JSON 응답 반환
		return c.Status(statusCode).JSON(fiber.Map{
			"error": message,
		})
	}
}
