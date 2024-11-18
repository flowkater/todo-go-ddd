package http

import (
	"log"
	"strconv"

	"github.com/flowkater/ddd-todo-app/internal/application/command"
	"github.com/flowkater/ddd-todo-app/internal/application/query"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	commandUsecase *command.TodoCommandUsecase
	queryUsecase   *query.TodoQueryUsecase
}

func NewTodoHandler(commandUsecase *command.TodoCommandUsecase, queryUsecase *query.TodoQueryUsecase) *TodoHandler {
	return &TodoHandler{
		commandUsecase: commandUsecase,
		queryUsecase:   queryUsecase,
	}
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var req CreateTodoRequest
	if err := c.BodyParser(&req); err != nil {
		log.Printf("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	log.Printf("received create todo request: %+v", req)
	id, err := h.commandUsecase.Usecase(c.Context(), command.CreateTodoCommand{
		Title:       req.Title,
		Description: req.Description,
	})

	if err != nil {
		log.Printf("failed to create todo: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Printf("todo created successfully with id: %d", id)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Todo created successfully",
		"id":     id,
	})
}

func (h *TodoHandler) GetTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Printf("failed to parse todo id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid todo ID",
		})
	}

	log.Printf("getting todo with id: %d", id)
	result, err := h.queryUsecase.Query(c.Context(), query.GetTodoQuery{ID: id})
	if err != nil {
		if err.Error() == "ent: todo not found" {
			log.Printf("todo not found with id: %d", id)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Todo not found",
			})
		}
		log.Printf("failed to get todo: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
