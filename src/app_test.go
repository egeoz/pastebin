package src_test

import (
	"context"
	_ "embed"
	"pastebin/src/controller"
	"pastebin/src/logger"
	"pastebin/src/repos"
	"testing"
)

//go:embed sql/sqlite/table_check.sql
var sqliteTableCheck string

//go:embed sql/sqlite/schema.sql
var sqliteDdl string

const localhost = "127.0.0.1"
const port = 32198

func TestRun(t *testing.T) {
	ctx := context.Background()
	logger.InitLogger("DEBUG")
	repository, err := repos.NewSqliteRepository(ctx, "::memory", sqliteDdl, sqliteTableCheck)
	if err != nil {
		t.Fatal(err)
	}

	controller := controller.NewController(repository, localhost, port)
	controller.Serve()
}
