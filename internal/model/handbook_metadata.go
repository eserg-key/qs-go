package model

import (
	"github.com/google/uuid"
	"time"
)

type HandbookMetaData struct {
	ID                     uuid.UUID  `json:"id"`
	Sort                   uint32     `json:"sort"`
	HandbookName           string     `json:"handbook_name"`
	TypeField              uint32     `json:"type_field"`
	HandbookFieldName      string     `json:"handbook_field_name"`
	HandbookFieldNameRUS   string     `json:"handbook_field_name_rus"`
	HandbookChildrenID     uint32     `json:"handbook_children_id"`
	HandbookChildrenColumn string     `json:"handbook_children_column"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              *time.Time `json:"updated_at"`
}

func NewHandbookMetaData(ID uuid.UUID, sort uint32, handbookName string, typeField uint32, handbookFieldName string, handbookFieldNameRUS string, HandbookChildrenID uint32, handbookChildrenColumn string, createdAt time.Time, updatedAt *time.Time) HandbookMetaData {
	return HandbookMetaData{
		ID:                     ID,
		Sort:                   sort,
		HandbookName:           handbookName,
		TypeField:              typeField,
		HandbookFieldName:      handbookFieldName,
		HandbookFieldNameRUS:   handbookFieldNameRUS,
		HandbookChildrenID:     HandbookChildrenID,
		HandbookChildrenColumn: handbookChildrenColumn,
		CreatedAt:              createdAt,
		UpdatedAt:              updatedAt,
	}
}

func GetFieldsName(metaData []HandbookMetaData) []string {
	fieldName := make([]string, 0, len(metaData))
	for _, val := range metaData {
		fieldName = append(fieldName, val.HandbookFieldName)
	}

	return fieldName
}

type CreateHandbookMetaData struct {
	ID                     uuid.UUID
	Sort                   uint32
	HandbookName           string
	TypeField              uint32
	HandbookFieldName      string
	HandbookFieldNameRUS   string
	HandbookChildrenID     uint32
	HandbookChildrenColumn string
	CreatedAt              time.Time
}

func NewCreateHandbookMetaData(
	ID uuid.UUID,
	sort uint32,
	handbookName string,
	typeField uint32,
	handbookFieldName string,
	handbookFieldNameRUS string,
	HandbookChildrenID uint32,
	handbookChildrenColumn string,
	createdAt time.Time,
) CreateHandbookMetaData {
	return CreateHandbookMetaData{
		ID:                     ID,
		Sort:                   sort,
		HandbookName:           handbookName,
		TypeField:              typeField,
		HandbookFieldName:      handbookFieldName,
		HandbookFieldNameRUS:   handbookFieldNameRUS,
		HandbookChildrenID:     HandbookChildrenID,
		HandbookChildrenColumn: handbookChildrenColumn,
		CreatedAt:              createdAt,
	}
}
