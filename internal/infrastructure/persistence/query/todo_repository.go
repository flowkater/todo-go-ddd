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

func (r *todoRepository) GetAll(ctx context.Context) ([]*entity.Todo, error) {
	todos, err := r.client.Todo.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}

	var result []*entity.Todo
	for _, t := range todos {
		result = append(result, &entity.Todo{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Completed:   t.Completed,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		})
	}

	return result, nil
}

// List implements query.TodoRepository.
// func (r *todoRepository) List(ctx context.Context, filter query.TodoFilter) ([]*entity.Todo, error) {
// 	panic("unimplemented")
// }

func NewTodoRepository(client *ent.Client) query.TodoRepository {
	return &todoRepository{client: client}
}

func (r *todoRepository) Summary(ctx context.Context) (total int, completed int, uncompleted int, err error) {
	// 전체 개수
	total, err = r.client.Todo.
		Query().
		Count(ctx)

	if err != nil {
		return 0, 0, 0, err
	}

	// 완료 개수
	completed, err = r.client.Todo.Query().Where(todo.Completed(true)).Count(ctx)
	if err != nil {
		return 0, 0, 0, err
	}

	// 미완료 개수
	uncompleted = total - completed

	return total, completed, uncompleted, nil
}
