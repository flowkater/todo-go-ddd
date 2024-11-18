package command

import (
	"context"

	"github.com/flowkater/ddd-todo-app/internal/domain/entity"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *entity.Todo) (int, error)
	Update(ctx context.Context, todo *entity.Todo) error
	Delete(ctx context.Context, id int) error
}
