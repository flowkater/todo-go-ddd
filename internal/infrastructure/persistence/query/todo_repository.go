package query

import (
	"context"

	"github.com/flowkater/ddd-todo-app/internal/domain/entity"
	"github.com/flowkater/ddd-todo-app/internal/domain/repository/query"
	"github.com/flowkater/ddd-todo-app/internal/infrastructure/persistence/ent"
	"github.com/flowkater/ddd-todo-app/internal/infrastructure/persistence/ent/todo"
)

type todoRepository struct {
	client *ent.Client
}

// GetById implements query.TodoRepository.
func (r *todoRepository) GetById(ctx context.Context, id int) (*entity.Todo, error) {
	t, err := r.client.Todo.
		Query().
		Where(todo.ID(id)).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return &entity.Todo{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}, nil
}

// List implements query.TodoRepository.
// func (r *todoRepository) List(ctx context.Context, filter query.TodoFilter) ([]*entity.Todo, error) {
// 	panic("unimplemented")
// }

func NewTodoRepository(client *ent.Client) query.TodoRepository {
	return &todoRepository{client: client}
}
