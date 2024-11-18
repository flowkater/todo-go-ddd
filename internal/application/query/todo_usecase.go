package query

import (
	"context"
	"time"

	"github.com/flowkater/ddd-todo-app/internal/domain/repository/query"
	"github.com/flowkater/ddd-todo-app/pkg/errors"
)

type TodoQueryUsecase struct {
	todoRepo query.TodoRepository
}

func NewTodoQueryUsecase(repo query.TodoRepository) *TodoQueryUsecase {
	return &TodoQueryUsecase{
		todoRepo: repo,
	}
}

func (u *TodoQueryUsecase) Query(ctx context.Context, q interface{}) (interface{}, error) {
	switch query := q.(type) {
	case GetTodoQuery:
		return u.getTodoQuery(ctx, query)
	default:
		return nil, errors.ErrUnknownQuery
	}
}

func (u *TodoQueryUsecase) getTodoQuery(ctx context.Context, q GetTodoQuery) (*TodoDTO, error) {
	todo, err := u.todoRepo.GetById(ctx, q.ID)
	if err != nil {
		return nil, err
	}

	return &TodoDTO{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   todo.UpdatedAt.Format(time.RFC3339),
	}, nil
}
