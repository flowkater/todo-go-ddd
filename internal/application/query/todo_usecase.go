package query

import (
	"context"

	"github.com/flowkater/ddd-todo-app/internal/domain/entity"
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
	case GetAllTodoQuery:
		return u.getAllTodoQuery(ctx, query)
	default:
		return nil, errors.ErrUnknownQuery
	}
}

func (u *TodoQueryUsecase) getTodoQuery(ctx context.Context, q GetTodoQuery) (*entity.Todo, error) {
	todo, err := u.todoRepo.GetById(ctx, q.ID)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (u *TodoQueryUsecase) getAllTodoQuery(ctx context.Context, q GetAllTodoQuery) ([]*entity.Todo, error) {
	todos, err := u.todoRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
