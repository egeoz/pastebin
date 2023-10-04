package repos

import (
	"context"
	"database/sql"
	dbHandler "pastebin/src/db/postgresql"
	"pastebin/src/dto"
	"pastebin/src/logger"
	"pastebin/src/repos/tools"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type PostgresqlRepository struct {
	db      *sql.DB
	handler *dbHandler.Queries
}

func NewPostgresqlRepository(ctx context.Context, dataSource string, ddl string, tableCheck string) (*PostgresqlRepository, error) {
	db, err := sql.Open("postgres", dataSource)

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

	return &PostgresqlRepository{db: db, handler: dbHandler.New(db)}, nil
}

func (repo PostgresqlRepository) GetEntry(ctx context.Context, uuid string) (dto.Entry, error) {
	entity, err := repo.handler.GetEntry(ctx, uuid)

	return dto.Entry{UUID: entity.Uuid, Title: entity.Title, Content: entity.Content, ContentType: entity.ContentType}, err
}

func (repo PostgresqlRepository) CreateEntry(ctx context.Context, entry dto.Entry) (string, error) {
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

func (repo PostgresqlRepository) DeleteEntry(ctx context.Context, uuid string) error {
	return repo.handler.DeleteEntry(ctx, uuid)
}
