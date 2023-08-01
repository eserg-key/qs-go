package policy

import (
	"context"
	"encoding/json"
	"fmt"
	"handbooks_backend/internal/model"
	"handbooks_backend/internal/policy/dto"
	"handbooks_backend/pkg/common/core/translation"
	"handbooks_backend/pkg/common/errors"
	"strings"
	"time"
)

var (
	errTypeFields = errors.New("Not type field")
)

func (p *Policy) CreateHandbook(ctx context.Context, input dto.CreateHandbookInput) (model.Handbook, error) {
	// Create Handbook
	createHandbook := model.NewCreateHandbook(
		input.HandbookName,
		translation.Translation(input.HandbookName),
		input.ProjectCode,
		time.Now(),
	)
	// Insert handbook
	// TODO обработка ошибок в коде и дублирования с выводом на фронт, когда будем эту часть реализовать
	handbook, err := p.service.Handbook.CreateHandbook(ctx, createHandbook)
	if err != nil {
		return model.Handbook{}, errors.Wrap(err, "ServiceHandbook.CreateHandbook")
	}

	// Create MetaData
	resultHandbookMetaData := make([]model.HandbookMetaData, 0, len(input.MetaData))
	for _, metadata := range input.MetaData {
		metaData := model.NewCreateHandbookMetaData(
			p.identity.GenerateUUID(),
			metadata.Sort,
			createHandbook.HandbookName,
			metadata.TypeField,
			translation.Translation(metadata.Name),
			metadata.Name,
			metadata.HandbookChildrenID,
			metadata.HandbookChildrenColumn,
			time.Now(),
		)
		// Insert metadata
		response, err := p.service.CreateHandbookMetadata(ctx, metaData)
		if err != nil {
			return model.Handbook{}, errors.Wrap(err, "ServiceHandbook.CreateHandbookMetadata")
		}

		resultHandbookMetaData = append(resultHandbookMetaData, response)
	}

	// Get type fields
	fields, err := p.service.GetHandbookTypeFields(ctx)
	if err != nil {
		return model.Handbook{}, errors.Wrap(err, "ServiceHandbook.GetHandbookTypeFields")
	}
	mapFields := make(map[uint32]model.TypeField)
	for _, f := range fields {
		mapFields[f.ID] = f
	}
	// Generate table
	sql := "CREATE TABLE " + createHandbook.HandbookName + " ("
	sql += "id UUID PRIMARY KEY,"
	for _, meta := range resultHandbookMetaData {
		field, ok := mapFields[meta.TypeField]
		if !ok {
			return model.Handbook{}, errors.Wrap(errTypeFields, "Policy.CreateHandbook")
		}

		fType := " TEXT"
		switch field.Type {
		case "int":
			fType = " INTEGER"
		case "text":
			fType = " VARCHAR(250)"
		case "select_city":
			fType = " JSON"
		case "select_handbook":
			fType = " JSON"
		}

		sql += meta.HandbookFieldName + fType + ","
	}
	sql += "created_at TIMESTAMPTZ NOT NULL,"
	sql += "updated_at TIMESTAMPTZ"
	sql += ")"
	// Create table
	err = p.service.Handbook.CreateTable(ctx, sql)
	if err != nil {
		return model.Handbook{}, errors.Wrap(err, "ServiceHandbook.CreateTable")
	}
	return handbook, nil
}

func (p *Policy) GetHandbooks(ctx context.Context, projectCode, search string) ([]model.Handbook, error) {
	return p.service.GetHandbooks(ctx, projectCode, search)
}

func (p *Policy) UpdateHandbook(ctx context.Context, id string, input dto.UpdateHandbookInput) (model.Handbook, error) {
	handbook, err := p.service.GetHandbook(ctx, id)
	if err != nil {
		return model.Handbook{}, errors.Wrap(err, "ServiceHandbook.GetHandbook")
	}

	timeNow := time.Now()
	updateFields := make(map[string]interface{})
	jsonData, _ := json.Marshal(input)
	_ = json.Unmarshal(jsonData, &updateFields)
	updateFields["updated_at"] = timeNow

	err = p.service.UpdateHandbook(ctx, id, updateFields)
	if err != nil {
		return model.Handbook{}, errors.Wrap(err, "ServiceHandbook.UpdateHandbook")
	}

	// Добавляем статически поля, доступные для обновления в output
	handbook.Name = input.Name
	handbook.UpdatedAt = &timeNow
	return handbook, nil
}

func (p *Policy) GetHandbook(ctx context.Context, id string, filter map[string]interface{}) (model.OutputRowsHandbook, error) {
	handbook, err := p.service.GetHandbook(ctx, id)
	if err != nil {
		return model.OutputRowsHandbook{}, errors.Wrap(err, "ServiceHandbook.GetHandbook")
	}
	return p.GetRowsHandbook(ctx, handbook, filter)
}

func (p *Policy) GetRowsHandbook(ctx context.Context, handbook model.Handbook, filter map[string]interface{}) (model.OutputRowsHandbook, error) {
	OutputRowsHandbook := model.OutputRowsHandbook{}
	meta, err := p.service.GetHandbookMetadata(ctx, handbook.HandbookName)
	if err != nil {
		return model.OutputRowsHandbook{}, errors.Wrap(err, "ServiceHandbook.GetHandbookMetadata")
	}

	metaData := model.GetFieldsName(meta)
	response, err := p.service.GetRowsHandbook(ctx, handbook.HandbookName, metaData, filter)
	if err != nil {
		return model.OutputRowsHandbook{}, errors.Wrap(err, "ServiceHandbook.GetRowsHandbook")
	}

	OutputRowsHandbook.ID = handbook.ID
	OutputRowsHandbook.Name = handbook.Name
	OutputRowsHandbook.MetaData = meta
	OutputRowsHandbook.Data = response

	return OutputRowsHandbook, nil
}

func (p *Policy) UpdateRowsHandbook(ctx context.Context, id string, input []map[string]interface{}) (model.OutputRowsHandbook, error) {
	handbook, err := p.service.GetHandbook(ctx, id)
	if err != nil {
		return model.OutputRowsHandbook{}, errors.Wrap(err, "ServiceHandbook.GetHandbook")
	}

	updateRows := make([]map[string]interface{}, 0, len(input))
	createRows := make([]map[string]interface{}, 0, len(input))
	idRows := make([]string, 0, len(input))
	for _, val := range input {
		if v, ok := val["id"]; ok && v != "" {
			val["updated_at"] = time.Now()
			updateRows = append(updateRows, val)
			idRows = append(idRows, fmt.Sprintf("'%s'", v))
		} else {
			val["id"] = p.identity.GenerateUUID()
			val["created_at"] = time.Now()
			idRows = append(idRows, fmt.Sprintf("'%s'", val["id"]))
			createRows = append(createRows, val)
		}
	}

	if len(updateRows) > 0 {
		err := p.service.UpdateRowsHandbook(ctx, handbook.HandbookName, updateRows)
		if err != nil {
			return model.OutputRowsHandbook{}, errors.Wrap(err, "ServiceHandbook.UpdateRowsHandbook")
		}
	}
	if len(createRows) > 0 {
		err = p.service.CreateRowsHandbook(ctx, handbook.HandbookName, createRows)
		if err != nil {
			return model.OutputRowsHandbook{}, errors.Wrap(err, "ServiceHandbook.CreateRowsHandbook")
		}
	}

	filter := make(map[string]interface{})
	filter["id"] = strings.Join(idRows, ",")
	return p.GetRowsHandbook(ctx, handbook, filter)
}
