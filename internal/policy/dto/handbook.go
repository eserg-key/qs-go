package dto

type CreateHandbookInput struct {
	HandbookName string `json:"handbook_name" binding:"required"`
	ProjectCode  string `json:"project_code" binding:"required"`
	MetaData     []struct {
		Sort                   uint32 `json:"sort" binding:"required"`
		Name                   string `json:"name" binding:"required"`
		TypeField              uint32 `json:"type_field" binding:"required"`
		HandbookChildrenID     uint32 `json:"handbook_children_id"`
		HandbookChildrenColumn string `json:"handbook_children_column"`
	} `json:"metadata"`
}

type UpdateHandbookInput struct {
	Name string `json:"name" binding:"required"`
}
