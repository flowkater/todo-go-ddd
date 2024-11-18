package http

import (
	"context"
	"fmt"

	"github.com/flowkater/ddd-todo-app/config"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app         *fiber.App
	config      *config.Config
	todoHandler *TodoHandler
}

func NewFiberApp(todoHandler *TodoHandler) *fiber.App {
	app := fiber.New()

	app.Post("/todos", todoHandler.CreateTodo)
	app.Get("/todos/:id", todoHandler.GetTodo)

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
