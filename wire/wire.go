//go:build wireinject

package wire

import (
	"fmt"

	"github.com/flowkater/ddd-todo-app/config"
	"github.com/flowkater/ddd-todo-app/internal/application/command"
	"github.com/flowkater/ddd-todo-app/internal/application/query"
	"github.com/flowkater/ddd-todo-app/internal/domain/service"
	cmdrepo "github.com/flowkater/ddd-todo-app/internal/infrastructure/persistence/command"
	"github.com/flowkater/ddd-todo-app/internal/infrastructure/persistence/database"
	"github.com/flowkater/ddd-todo-app/internal/infrastructure/persistence/ent"
	queryrepo "github.com/flowkater/ddd-todo-app/internal/infrastructure/persistence/query"
	"github.com/flowkater/ddd-todo-app/internal/interfaces/http"
	"github.com/google/wire"
)

func provideEntOptions(cfg *config.Config) ([]ent.Option, error) {
	drv, err := database.NewPostgresDriver(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres driver: %v", err)
	}

	return []ent.Option{ent.Driver(drv)}, nil
}

var infrastructureSet = wire.NewSet(
	provideEntOptions,
	ent.NewClient,
	cmdrepo.NewTodoRepository,
	queryrepo.NewTodoRepository,
)

var serviceSet = wire.NewSet(
	service.NewTodoService,
)

var handlerSet = wire.NewSet(
	command.NewTodoCommandUsecase,
	query.NewTodoQueryUsecase,
	http.NewTodoHandler,
)

var serverSet = wire.NewSet(
	http.NewFiberApp,
	http.NewServer,
)

func InitializeServer(cfg *config.Config) (*http.Server, error) {
	panic(wire.Build(
		infrastructureSet,
		serviceSet,
		handlerSet,
		serverSet,
	))
}
