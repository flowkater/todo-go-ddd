package command

import (
	"context"
	"log"

	"github.com/flowkater/ddd-todo-app/internal/domain/entity"
	"github.com/flowkater/ddd-todo-app/internal/domain/repository/command"
	"github.com/flowkater/ddd-todo-app/internal/infrastructure/persistence/ent"
)

type todoRepository struct {
	client *ent.Client
}

func NewTodoRepository(client *ent.Client) command.TodoRepository {
	return &todoRepository{
		client: client,
	}
}

func (r *todoRepository) Create(ctx context.Context, todo *entity.Todo) (int, error) {
	created, err := r.client.Todo.
		Create().
		SetTitle(todo.Title).
		SetDescription(todo.Description).
		SetCompleted(todo.Completed).
		Save(ctx)

	if err != nil {
		log.Printf("failed to create todo: %v", err)
		return 0, err
	}

	log.Printf("created todo with id: %d", created.ID)
	return created.ID, nil
}

func (r *todoRepository) Update(ctx context.Context, todo *entity.Todo) error {
	return r.client.Todo.
		UpdateOneID(todo.ID).
		SetTitle(todo.Title).
		SetDescription(todo.Description).
		SetCompleted(todo.Completed).
		Exec(ctx)
}

func (r *todoRepository) Delete(ctx context.Context, id int) error {
	return r.client.Todo.DeleteOneID(id).Exec(ctx)
}
