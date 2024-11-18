package database

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/flowkater/ddd-todo-app/config"
	"github.com/flowkater/ddd-todo-app/internal/infrastructure/persistence/ent"
	_ "github.com/lib/pq"
)

func NewPostgresDriver(cfg *config.Config) (*sql.Driver, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
	)

	drv, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}

	return drv, nil
}

func RunMigration(client *ent.Client) error {
	if err := client.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}
	return nil
}
