package query

import (
	"context"

	"github.com/flowkater/ddd-todo-app/internal/domain/entity"
)

type TodoRepository interface {
	GetById(ctx context.Context, id int) (*entity.Todo, error)
	GetAll(ctx context.Context) ([]*entity.Todo, error)
	// List(ctx context.Context, filter TodoFilter) ([]*entity.Todo, error)
	Summary(ctx context.Context) (total int, completed int, uncompleted int, err error)
}

type TodoFilter struct {
	Completed *bool
	Search    string
	Limit     int
	Offset    int
}
