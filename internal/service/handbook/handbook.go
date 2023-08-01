package handbook

import (
	"context"
	"handbooks_backend/internal/model"
	"handbooks_backend/internal/repository"
)

type ServiceHandbook struct {
	repo *repository.Repository
}

func NewServiceHandbook(repo *repository.Repository) *ServiceHandbook {
	return &ServiceHandbook{repo: repo}
}

func (s *ServiceHandbook) CreateHandbook(ctx context.Context, req model.CreateHandbook) (model.Handbook, error) {
	id, err := s.repo.CreateHandbook(ctx, req)
	if err != nil {
		return model.Handbook{}, err
	}

	return model.NewHandbook(
		id,
		req.Name,
		req.HandbookName,
		req.ProjectCode,
		req.CreatedAt,
		nil,
	), nil
}

func (s *ServiceHandbook) CreateHandbookMetadata(ctx context.Context, req model.CreateHandbookMetaData) (model.HandbookMetaData, error) {
	id, err := s.repo.CreateHandbookMetadata(ctx, req)
	if err != nil {
		return model.HandbookMetaData{}, err
	}

	return model.NewHandbookMetaData(
		id,
		req.Sort,
		req.HandbookName,
		req.TypeField,
		req.HandbookFieldName,
		req.HandbookFieldNameRUS,
		req.HandbookChildrenID,
		req.HandbookChildrenColumn,
		req.CreatedAt,
		nil,
	), nil
}

func (s *ServiceHandbook) GetHandbookMetadata(ctx context.Context, handbookName string) ([]model.HandbookMetaData, error) {
	return s.repo.GetHandbookMetadata(ctx, handbookName)
}

func (s *ServiceHandbook) DeleteRowTable(ctx context.Context, id, table string) error {
	err := s.repo.DeleteRowTable(ctx, id, table)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceHandbook) CreateTable(ctx context.Context, sql string) error {
	err := s.repo.CreateTable(ctx, sql)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceHandbook) GetHandbookTypeFields(ctx context.Context) ([]model.TypeField, error) {
	fields, err := s.repo.GetHandbookTypeFields(ctx)
	if err != nil {
		return []model.TypeField{}, err
	}
	return fields, nil
}

func (s *ServiceHandbook) GetHandbooks(ctx context.Context, projectCode, search string) ([]model.Handbook, error) {
	return s.repo.GetHandbooks(ctx, projectCode, search)
}

func (s *ServiceHandbook) GetHandbook(ctx context.Context, id string) (model.Handbook, error) {
	return s.repo.GetHandbook(ctx, id)
}

func (s *ServiceHandbook) UpdateHandbook(ctx context.Context, id string, fields map[string]interface{}) error {
	return s.repo.UpdateHandbook(ctx, id, fields)
}

func (s *ServiceHandbook) CreateRowsHandbook(ctx context.Context, handbookName string, req []map[string]interface{}) error {
	for _, val := range req {
		err := s.repo.CreateRowHandbook(ctx, handbookName, val)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ServiceHandbook) UpdateRowsHandbook(ctx context.Context, handbookName string, fields []map[string]interface{}) error {
	for _, val := range fields {
		err := s.repo.UpdateRowsHandbook(ctx, handbookName, val)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ServiceHandbook) GetRowsHandbook(ctx context.Context, projectCode string, metaData []string, filter map[string]interface{}) ([]map[string]interface{}, error) {
	return s.repo.GetRowsHandbook(ctx, projectCode, metaData, filter)
}
