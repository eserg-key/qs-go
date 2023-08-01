package handbook

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"handbooks_backend/internal/dal"
	postgre "handbooks_backend/internal/dal/postgres"
	"handbooks_backend/internal/model"
	"handbooks_backend/pkg/postgres"
	"strings"
)

type RepositoryHandbook struct {
	queryBuilder squirrel.StatementBuilderType
	client       postgres.Client
}

func NewRepositoryHandbook(client postgres.Client) *RepositoryHandbook {
	return &RepositoryHandbook{
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		client:       client,
	}
}

func (repo *RepositoryHandbook) CreateHandbook(ctx context.Context, req model.CreateHandbook) (uint32, error) {
	var id uint32
	sql, args, err := repo.queryBuilder.
		Insert(postgre.HandbookTable).
		Columns(
			"name",
			"handbook_name",
			"project_code",
			"created_at",
		).
		Values(
			req.Name,
			req.HandbookName,
			req.ProjectCode,
			req.CreatedAt,
		).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return 0, err
	}

	execErr := repo.client.QueryRow(ctx, sql, args...).Scan(&id)
	if execErr != nil {
		return 0, execErr
	}

	return id, nil
}

func (repo *RepositoryHandbook) CreateHandbookMetadata(ctx context.Context, req model.CreateHandbookMetaData) (uuid.UUID, error) {
	sql, args, err := repo.queryBuilder.
		Insert(postgre.HandbookMetadataTable).
		Columns(
			"id",
			"sort",
			"handbook_name",
			"handbook_type_field_id",
			"handbook_field_name",
			"handbook_field_rus_name",
			"handbook_children_id",
			"handbook_children_column",
			"created_at",
		).
		Values(
			req.ID,
			req.Sort,
			req.HandbookName,
			req.TypeField,
			req.HandbookFieldName,
			req.HandbookFieldNameRUS,
			req.HandbookChildrenID,
			req.HandbookChildrenColumn,
			req.CreatedAt,
		).
		ToSql()
	if err != nil {
		return req.ID, err
	}

	_, execErr := repo.client.Exec(ctx, sql, args...)
	if execErr != nil {
		return req.ID, execErr
	}
	return req.ID, nil
}

func (repo *RepositoryHandbook) GetHandbookMetadata(ctx context.Context, handbookName string) ([]model.HandbookMetaData, error) {
	query := repo.queryBuilder.
		Select(
			"id",
			"sort",
			"handbook_name",
			"handbook_type_field_id",
			"handbook_field_name",
			"handbook_field_rus_name",
			"handbook_children_id",
			"handbook_children_column",
			"created_at",
			"updated_at",
		)
	query = query.From(postgre.HandbookMetadataTable).
		Where(squirrel.Eq{"handbook_name": handbookName}).OrderByClause("sort")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := repo.client.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	entities := make([]model.HandbookMetaData, rows.CommandTag().RowsAffected())

	for rows.Next() {
		var e model.HandbookMetaData
		if err = rows.Scan(
			&e.ID,
			&e.Sort,
			&e.HandbookName,
			&e.TypeField,
			&e.HandbookFieldName,
			&e.HandbookFieldNameRUS,
			&e.HandbookChildrenID,
			&e.HandbookChildrenColumn,
			&e.CreatedAt,
			&e.UpdatedAt,
		); err != nil {
			return nil, err
		}

		entities = append(entities, e)
	}

	return entities, nil
}

func (repo *RepositoryHandbook) DeleteRowTable(ctx context.Context, id, table string) error {
	sql, args, err := repo.queryBuilder.
		Delete(table).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	cmd, execErr := repo.client.Exec(ctx, sql, args...)
	if execErr != nil {
		return execErr
	}
	if cmd.RowsAffected() == 0 {
		return dal.ErrNothingInserted
	}

	return nil
}

func (repo *RepositoryHandbook) CreateTable(ctx context.Context, sql string) error {
	_, execErr := repo.client.Exec(ctx, sql)
	if execErr != nil {
		return execErr
	}

	return nil
}

func (repo *RepositoryHandbook) GetHandbookTypeFields(ctx context.Context) ([]model.TypeField, error) {
	statement := repo.queryBuilder.
		Select(
			"id",
			"name",
			"type",
		).
		From(postgre.HandbookTypeFields)

	query, args, err := statement.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := repo.client.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	entities := make([]model.TypeField, rows.CommandTag().RowsAffected())

	for rows.Next() {
		var e model.TypeField
		if err = rows.Scan(
			&e.ID,
			&e.Name,
			&e.Type,
		); err != nil {
			return nil, err
		}

		entities = append(entities, e)
	}

	return entities, nil
}

func (repo *RepositoryHandbook) GetHandbooks(ctx context.Context, projectCode, search string) ([]model.Handbook, error) {
	query := repo.queryBuilder.
		Select(
			"id",
			"name",
			"handbook_name",
			"project_code",
			"created_at",
			"updated_at",
		)
	query = query.From(postgre.HandbookTable)
	if projectCode != "" {
		query = query.Where(squirrel.Eq{"project_code": projectCode})
	}
	if search != "" {
		query = query.Where(squirrel.ILike{"name": fmt.Sprintf("%%%s%%", search)})
	}
	sql, args, err := query.OrderBy("created_at ASC").ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := repo.client.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	entities := make([]model.Handbook, rows.CommandTag().RowsAffected())

	for rows.Next() {
		var e model.Handbook
		if err = rows.Scan(
			&e.ID,
			&e.Name,
			&e.HandbookName,
			&e.ProjectCode,
			&e.CreatedAt,
			&e.UpdatedAt,
		); err != nil {
			return nil, err
		}

		entities = append(entities, e)
	}

	return entities, nil
}

func (repo *RepositoryHandbook) GetHandbook(ctx context.Context, id string) (model.Handbook, error) {
	sql, args, err := repo.queryBuilder.
		Select(
			"id",
			"name",
			"handbook_name",
			"project_code",
			"created_at",
			"updated_at",
		).
		From(postgre.HandbookTable).
		Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return model.Handbook{}, err
	}

	var h model.Handbook
	err = repo.client.QueryRow(ctx, sql, args...).Scan(
		&h.ID,
		&h.Name,
		&h.HandbookName,
		&h.ProjectCode,
		&h.CreatedAt,
		&h.UpdatedAt,
	)
	if err != nil {
		return model.Handbook{}, err
	}

	return h, nil
}

func (repo *RepositoryHandbook) UpdateHandbook(ctx context.Context, id string, fields map[string]interface{}) error {
	sql, args, err := repo.queryBuilder.
		Update(postgre.HandbookTable).
		SetMap(fields).
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	if exec, err := repo.client.Exec(ctx, sql, args...); err != nil {
		return err
	} else if exec.RowsAffected() == 0 || !exec.Update() {
		return dal.ErrNothingInserted
	}

	return nil
}

func (repo *RepositoryHandbook) CreateRowHandbook(ctx context.Context, handbookName string, req map[string]interface{}) error {
	keys := make([]string, 0, len(req))
	values := make([]interface{}, 0, len(req))
	for i, val := range req {
		keys = append(keys, i)
		values = append(values, val)
	}
	sql, args, err := repo.queryBuilder.
		Insert("public." + handbookName).
		Columns(keys...).
		Values(values...).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := repo.client.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

func (repo *RepositoryHandbook) UpdateRowsHandbook(ctx context.Context, handbookName string, fields map[string]interface{}) error {
	id := fields["id"]
	delete(fields, "id")
	sql, args, err := repo.queryBuilder.
		Update("public." + handbookName).
		SetMap(fields).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	if exec, err := repo.client.Exec(ctx, sql, args...); err != nil {
		return err
	} else if exec.RowsAffected() == 0 || !exec.Update() {
		return dal.ErrNothingInserted
	}

	return nil
}

func (repo *RepositoryHandbook) GetRowsHandbook(ctx context.Context, projectCode string, metaData []string, filter map[string]interface{}) ([]map[string]interface{}, error) {
	metaData = append(metaData, "id")

	q := "SELECT "
	for _, meta := range metaData {
		q += meta + ","
	}
	q = strings.TrimRight(q, ",")
	q += " FROM " + "public." + projectCode + " "

	if v, ok := filter["search"]; ok {
		q += "WHERE "
		for _, meta := range metaData {
			if meta == "id" {
				continue
			}
			q += meta + "::text ILIKE '" + fmt.Sprintf("%%%s%%", v) + "' OR "
		}
	}
	q = strings.TrimRight(q, " OR ")

	if v, ok := filter["id"]; ok {
		if _, ok = filter["search"]; ok {
			q += " AND "
		} else {
			q += " WHERE "
		}
		q += "id IN(" + v.(string) + ")"
	}

	q += " ORDER BY created_at ASC"

	rows, err := repo.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entities := make([]map[string]interface{}, 0)
	for rows.Next() {
		values := make([]interface{}, len(metaData))
		idType := make([]uuid.UUID, len(metaData))
		valuePointers := make([]interface{}, len(metaData))
		for i, v := range metaData {
			if v == "id" {
				valuePointers[i] = &idType[i]
			} else {
				valuePointers[i] = &values[i]
			}
			//valuePointers[i] = &values[i]
		}
		if err := rows.Scan(valuePointers...); err != nil {
			return nil, err
		}
		entity := make(map[string]interface{})
		for i, column := range metaData {
			if column == "id" {
				entity[column] = idType[i]
			} else {
				entity[column] = values[i]
			}
		}
		entities = append(entities, entity)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return entities, nil
}
