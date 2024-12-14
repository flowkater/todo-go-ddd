package middleware

import (
	"github.com/flowkater/ddd-todo-app/internal/interfaces/http/dto"
	"github.com/flowkater/ddd-todo-app/internal/interfaces/http/errors"
	"github.com/flowkater/ddd-todo-app/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

// ValidateRequest returns a middleware function that validates the request body
func ValidateRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// POST, PUT, PATCH 요청에 대해서만 검증
		if c.Method() != fiber.MethodPost && c.Method() != fiber.MethodPut && c.Method() != fiber.MethodPatch {
			return c.Next()
		}

		// 경로에 따라 다른 DTO 사용
		switch c.Path() {
		case "/todos":
			if c.Method() == fiber.MethodPost {
				var req dto.CreateTodoRequest
				if err := c.BodyParser(&req); err != nil {
					return errors.NewHTTPError(fiber.StatusBadRequest, "Invalid request body", err)
				}
				if err := req.Validate(); err != nil {
					// ValidationErrors 타입이면 그대로 반환
					if validationErrors, ok := err.(*validator.ValidationErrors); ok {
						return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
					}
					return errors.NewHTTPError(fiber.StatusBadRequest, err.Error(), err)
				}
				// 검증된 요청을 locals에 저장
				c.Locals("request", req)
			}
		}

		return c.Next()
	}
}
