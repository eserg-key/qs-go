package repository

import (
	"context"
	"github.com/google/uuid"
	"handbooks_backend/internal/model"
	"handbooks_backend/internal/repository/handbook"
	"handbooks_backend/pkg/postgres"
)

type Repository struct {
	Handbook
}

func NewRepository(db postgres.Client) *Repository {
	return &Repository{
		Handbook: handbook.NewRepositoryHandbook(db),
	}
}

type Handbook interface {
	CreateHandbook(context.Context, model.CreateHandbook) (uint32, error)
	CreateHandbookMetadata(context.Context, model.CreateHandbookMetaData) (uuid.UUID, error)
	GetHandbookMetadata(context.Context, string) ([]model.HandbookMetaData, error)
	DeleteRowTable(ctx context.Context, id, table string) error
	CreateTable(context.Context, string) error
	GetHandbookTypeFields(context.Context) ([]model.TypeField, error)
	GetHandbooks(context.Context, string, string) ([]model.Handbook, error)
	GetHandbook(context.Context, string) (model.Handbook, error)
	UpdateHandbook(context.Context, string, map[string]interface{}) error
	CreateRowHandbook(context.Context, string, map[string]interface{}) error
	UpdateRowsHandbook(context.Context, string, map[string]interface{}) error
	GetRowsHandbook(context.Context, string, []string, map[string]interface{}) ([]map[string]interface{}, error)
}
