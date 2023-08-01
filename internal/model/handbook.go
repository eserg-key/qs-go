package model

import "time"

type Handbook struct {
	ID           uint32     `json:"id"`
	Name         string     `json:"name"`
	HandbookName string     `json:"handbook_name"`
	ProjectCode  string     `json:"project_code"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

func NewHandbook(
	ID uint32,
	name string,
	handbookName string,
	projectCode string,
	createdAt time.Time,
	updatedAt *time.Time,
) Handbook {
	return Handbook{
		ID:           ID,
		Name:         name,
		HandbookName: handbookName,
		ProjectCode:  projectCode,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}

type OutputRowsHandbook struct {
	ID       uint32                   `json:"id"`
	Name     string                   `json:"name"`
	MetaData []HandbookMetaData       `json:"metadata"`
	Data     []map[string]interface{} `json:"data"`
}

type CreateHandbook struct {
	Name         string
	HandbookName string
	ProjectCode  string
	CreatedAt    time.Time
}

func NewCreateHandbook(
	name string,
	handbookName string,
	projectCode string,
	createdAt time.Time,
) CreateHandbook {
	return CreateHandbook{
		Name:         name,
		HandbookName: handbookName,
		ProjectCode:  projectCode,
		CreatedAt:    createdAt,
	}
}
