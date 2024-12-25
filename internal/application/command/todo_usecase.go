package command

import (
	"context"
	"log"

	"github.com/flowkater/ddd-todo-app/internal/domain/entity"
	"github.com/flowkater/ddd-todo-app/internal/domain/repository/command"
	"github.com/flowkater/ddd-todo-app/internal/domain/service"
	"github.com/flowkater/ddd-todo-app/pkg/errors"
)

type TodoCommandUsecase struct {
	todoService    service.TodoService
	todoRepository command.TodoRepository
}

func NewTodoCommandUsecase(todoService service.TodoService, todoRepository command.TodoRepository) *TodoCommandUsecase {
	return &TodoCommandUsecase{
		todoService:    todoService,
		todoRepository: todoRepository,
	}
}

func (u *TodoCommandUsecase) Usecase(ctx context.Context, cmd interface{}) (interface{}, error) {
	switch c := cmd.(type) {
	case CreateTodoCommand:
		return u.createUsecase(ctx, c)
	default:
		return nil, errors.ErrUnknownCommand
	}
}

func (u *TodoCommandUsecase) createUsecase(ctx context.Context, cmd CreateTodoCommand) (int, error) {
	log.Printf("creating todo with title: %s", cmd.Title)
	todo := u.todoService.New(&entity.TodoCreation{
		Title:       cmd.Title,
		Description: cmd.Description,
	})
	id, err := u.todoRepository.Create(ctx, todo)
	if err != nil {
		log.Printf("failed to create todo in usecase: %v", err)
		return 0, err
	}
	log.Printf("todo created successfully with id: %d", id)
	return id, nil
}

func (u *TodoCommandUsecase) DeleteTodo(ctx context.Context, cmd DeleteTodoCommand) error {
	log.Printf("deleting todo with id: %d", cmd.ID)
	return u.todoRepository.Delete(ctx, cmd.ID)
}
