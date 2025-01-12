package http

import (
	"strconv"

	"github.com/flowkater/ddd-todo-app/internal/application/command"
	"github.com/flowkater/ddd-todo-app/internal/application/query"
	"github.com/flowkater/ddd-todo-app/internal/domain/entity"
	"github.com/flowkater/ddd-todo-app/internal/interfaces/http/dto"
	"github.com/flowkater/ddd-todo-app/internal/interfaces/http/errors"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type TodoHandler struct {
	commandUsecase *command.TodoCommandUsecase
	queryUsecase   *query.TodoQueryUsecase
	logger         *zap.Logger
}

func NewTodoHandler(commandUsecase *command.TodoCommandUsecase, queryUsecase *query.TodoQueryUsecase) *TodoHandler {
	logger, _ := zap.NewProduction()
	return &TodoHandler{
		commandUsecase: commandUsecase,
		queryUsecase:   queryUsecase,
		logger:         logger,
	}
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var req dto.CreateTodoRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.NewHTTPError(fiber.StatusBadRequest, "Invalid request body", err)
	}

	h.logger.Info("received create todo request",
		zap.String("title", req.Title),
		zap.String("description", req.Description),
	)

	todoCreation := req.ToEntity()
	id, err := h.commandUsecase.Usecase(c.Context(), command.CreateTodoCommand{
		Title:       todoCreation.Title,
		Description: todoCreation.Description,
	})

	if err != nil {
		return errors.NewHTTPError(fiber.StatusInternalServerError, "Failed to create todo", err)
	}

	h.logger.Info("todo created successfully",
		zap.Int("id", id.(int)),
	)

	response := dto.CreateTodoResponse{
		Message: "Todo created successfully",
		ID:      id,
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *TodoHandler) GetTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errors.NewHTTPError(fiber.StatusBadRequest, "Invalid todo ID", err)
	}

	h.logger.Info("getting todo",
		zap.Int("id", id),
	)

	result, err := h.queryUsecase.Query(c.Context(), query.GetTodoQuery{ID: id})
	if err != nil {
		if errors.IsNotFound(err) {
			return errors.NewHTTPError(fiber.StatusNotFound, "Todo not found", err)
		}
		return errors.NewHTTPError(fiber.StatusInternalServerError, "Failed to get todo", err)
	}

	todo := result.(*entity.Todo)
	response := dto.TodoResponseFromEntity(todo)
	return c.Status(fiber.StatusOK).JSON(response)
}
