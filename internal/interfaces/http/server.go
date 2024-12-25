package http

import (
	"context"
	"fmt"

	"github.com/flowkater/ddd-todo-app/config"
	"github.com/flowkater/ddd-todo-app/internal/application/command"
	"github.com/flowkater/ddd-todo-app/internal/application/query"
	"github.com/flowkater/ddd-todo-app/internal/interfaces/http/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

type Server struct {
	app         *fiber.App
	config      *config.Config
	todoHandler *TodoHandler
}

func NewFiberApp(commandUsecase *command.TodoCommandUsecase, queryUsecase *query.TodoQueryUsecase) *fiber.App {
	// 에러 처리 설정
	app := fiber.New(fiber.Config{
		// 에러 핸들러 설정
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// 이미 상태 코드가 설정된 경우
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// 에러 응답
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
		// 서버가 패닉 상황에서도 종료되지 않도록 설정
		DisableStartupMessage: false,
	})

	// 로거 설정
	zapLogger, _ := zap.NewProduction()

	// 미들웨어 설정
	app.Use(recover.New()) // 패닉 복구 미들웨어
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(middleware.ErrorMiddleware(zapLogger))
	app.Use(middleware.ValidateRequest())

	// 핸들러 설정
	todoHandler := NewTodoHandler(commandUsecase, queryUsecase)

	// 라우트 설정
	app.Post("/todos", todoHandler.CreateTodo)
	app.Get("/todos/list", todoHandler.GetAllTodos)
	app.Get("/todos/:id", todoHandler.GetTodo)
	app.Delete("/todos/:id", todoHandler.DeleteTodo)

	return app
}

func NewServer(app *fiber.App, config *config.Config, todoHandler *TodoHandler) *Server {
	return &Server{
		app:         app,
		config:      config,
		todoHandler: todoHandler,
	}
}

func (s *Server) Start() error {
	return s.app.Listen(fmt.Sprintf(":%s", s.config.Server.Port))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}
