package service

import (
	"context"
	"github.com/redis/go-redis/v9"
	"handbooks_backend/internal/model"
	"handbooks_backend/internal/repository"
	"handbooks_backend/internal/service/handbook"
)

type Service struct {
	Handbook
	redis *redis.Client
}

func NewService(repo *repository.Repository, redis *redis.Client) *Service {
	return &Service{
		Handbook: handbook.NewServiceHandbook(repo),
		redis:    redis,
	}
}

type Handbook interface {
	CreateHandbook(context.Context, model.CreateHandbook) (model.Handbook, error)
	CreateHandbookMetadata(context.Context, model.CreateHandbookMetaData) (model.HandbookMetaData, error)
	GetHandbookMetadata(context.Context, string) ([]model.HandbookMetaData, error)
	DeleteRowTable(ctx context.Context, id, table string) error
	CreateTable(context.Context, string) error
	GetHandbookTypeFields(context.Context) ([]model.TypeField, error)
	GetHandbooks(context.Context, string, string) ([]model.Handbook, error)
	GetHandbook(context.Context, string) (model.Handbook, error)
	UpdateHandbook(context.Context, string, map[string]interface{}) error
	CreateRowsHandbook(context.Context, string, []map[string]interface{}) error
	UpdateRowsHandbook(context.Context, string, []map[string]interface{}) error
	GetRowsHandbook(context.Context, string, []string, map[string]interface{}) ([]map[string]interface{}, error)
}
