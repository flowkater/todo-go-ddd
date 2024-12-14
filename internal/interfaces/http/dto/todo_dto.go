package dto

import (
	"github.com/flowkater/ddd-todo-app/internal/domain/entity"
	"github.com/flowkater/ddd-todo-app/pkg/validator"
)

// CreateTodoRequest represents the request for creating a todo
type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"required,min=1,max=500"`
}

// Validate validates the CreateTodoRequest
func (r *CreateTodoRequest) Validate() error {
	return validator.Validate(r)
}

func (r *CreateTodoRequest) ToEntity() *entity.Todo {
	return &entity.Todo{
		Title:       r.Title,
		Description: r.Description,
	}
}

type CreateTodoResponse struct {
	Message string      `json:"message"`
	ID      interface{} `json:"id"`
}

type TodoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func TodoResponseFromEntity(todo *entity.Todo) *TodoResponse {
	return &TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt.String(),
		UpdatedAt:   todo.UpdatedAt.String(),
	}
}
