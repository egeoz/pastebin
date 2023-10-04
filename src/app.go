package src

import (
	"context"
	_ "embed"
	"fmt"
	"pastebin/src/config"
	"pastebin/src/controller"
	"pastebin/src/repos"
)

//go:embed sql/sqlite/table_check.sql
var sqliteTableCheck string

//go:embed sql/sqlite/schema.sql
var sqliteDdl string

//go:embed sql/postgresql/table_check.sql
var postgresqlTableCheck string

//go:embed sql/postgresql/schema.sql
var postgresqlDdl string

func Run() error {
	ctx := context.Background()
	cfg := config.LoadConfig()
	var repository repos.Repository
	var err error

	switch cfg.DatabaseType {
	case "sqlite":
		repository, err = repos.NewSqliteRepository(ctx, cfg.Database, sqliteDdl, sqliteTableCheck)
	case "postgresql":
		repository, err = repos.NewPostgresqlRepository(ctx, cfg.Database, postgresqlDdl, postgresqlTableCheck)
	default:
		return fmt.Errorf("invalid DatabaseType value: %s", cfg.DatabaseType)
	}

	if err != nil {
		return err
	}

	controller := controller.NewController(repository, cfg.Host, cfg.Port)
	controller.Serve()

	return nil
}
