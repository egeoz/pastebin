package repos

import (
	"context"
	"database/sql"
	dbHandler "pastebin/src/db/sqlite"
	"pastebin/src/dto"
	"pastebin/src/logger"
	"pastebin/src/repos/tools"
	"strconv"
	"time"

	_ "github.com/glebarez/go-sqlite"
)

type SqliteRepository struct {
	db      *sql.DB
	handler *dbHandler.Queries
}

func NewSqliteRepository(ctx context.Context, dataSource string, ddl string, tableCheck string) (*SqliteRepository, error) {
	db, err := sql.Open("sqlite", dataSource)

	if err != nil {
		return nil, err
	}

	tableExists, err := tools.CheckIfTableExists(db, ctx, tableCheck)
	if err != nil {
		return nil, err
	} else if !tableExists {
		logger.Log.Debug("Creating new table")
		if err = tools.CreateTable(db, ctx, ddl); err != nil {
			return nil, err
		}
	}

	return &SqliteRepository{db: db, handler: dbHandler.New(db)}, nil
}

func (repo SqliteRepository) GetEntry(ctx context.Context, uuid string) (dto.Entry, error) {
	entity, err := repo.handler.GetEntry(ctx, uuid)

	return dto.Entry{UUID: entity.Uuid, Title: entity.Title, Content: entity.Content, ContentType: entity.ContentType}, err
}

func (repo SqliteRepository) CreateEntry(ctx context.Context, entry dto.Entry) (string, error) {
	res, err := repo.handler.CreateEntry(ctx, dbHandler.CreateEntryParams{
		Uuid:        entry.UUID,
		Title:       entry.Title,
		Content:     entry.Content,
		ContentType: entry.ContentType,
		IsEncrypted: strconv.FormatBool(*entry.Encrypted),
		InsertDate:  time.Now().Format("2006-01-02 15:04:05.000"),
	})
	return res.Uuid, err
}

func (repo SqliteRepository) DeleteEntry(ctx context.Context, uuid string) error {
	return repo.handler.DeleteEntry(ctx, uuid)
}
