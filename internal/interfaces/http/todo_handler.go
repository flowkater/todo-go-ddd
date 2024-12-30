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

func (h *TodoHandler) GetAllTodos(c *fiber.Ctx) error {
	h.logger.Info("getting all todos")

	result, err := h.queryUsecase.Query(c.Context(), query.GetAllTodoQuery{})
	if err != nil {
		return errors.NewHTTPError(fiber.StatusInternalServerError, "Failed to get todos", err)
	}

	todos := result.([]*entity.Todo)
	response := dto.TodosResponseFromEntity(todos)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid todo id")
	}

	cmd := command.DeleteTodoCommand{ID: id}
	if err := h.commandUsecase.DeleteTodo(c.Context(), cmd); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid todo id")
	}

	var req dto.UpdateTodoRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.NewHTTPError(fiber.StatusBadRequest, "Invalid request body", err)
	}

	h.logger.Info("received update todo request",
		zap.Int("id", id),
		zap.String("title", req.Title),
		zap.String("description", req.Description),
	)

	cmd := command.UpdateTodoCommand{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
	}
	if err := h.commandUsecase.UpdateTodo(c.Context(), cmd); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *TodoHandler) ToggleTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid todo id")
	}

	cmd := command.ToggleTodoCommand{ID: id}
	if err := h.commandUsecase.ToggleTodo(c.Context(), cmd); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *TodoHandler) Summary(c *fiber.Ctx) error {
	total, completed, uncompleted, err := h.queryUsecase.Summary(c.Context())
	if err != nil {
		return errors.NewHTTPError(fiber.StatusInternalServerError, "Failed to get todo summary", err)
	}

	response := dto.TodoSummaryResponseFromEntity(total, completed, uncompleted)
	return c.Status(fiber.StatusOK).JSON(response)
}
