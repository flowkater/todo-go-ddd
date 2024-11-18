package service

import (
	"time"

	"github.com/flowkater/ddd-todo-app/internal/domain/entity"
)

type TodoService interface {
	New(creation *entity.TodoCreation) *entity.Todo
	Edit(todo *entity.Todo, title, description string) *entity.Todo
}

type todoService struct{}

func NewTodoService() TodoService {
	return &todoService{}
}

func (t *todoService) New(creation *entity.TodoCreation) *entity.Todo {
	now := time.Now()
	return &entity.Todo{
		Title:       creation.Title,
		Description: creation.Description,
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (t *todoService) Edit(todo *entity.Todo, title, description string) *entity.Todo {
	todo.Title = title
	todo.Description = description
	todo.UpdatedAt = time.Now()
	return todo
}
