package repos

import (
	"context"
	"pastebin/src/dto"
)

type Repository interface {
	GetEntry(ctx context.Context, uuid string) (dto.Entry, error)
	CreateEntry(ctx context.Context, entry dto.Entry) (string, error)
	DeleteEntry(ctx context.Context, uuid string) error
}
