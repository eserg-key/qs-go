package policy

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"handbooks_backend/internal/policy/dto"
)

func (p *Policy) AllProject(ctx *gin.Context, projectsCTX any) []dto.AllProjectOutput {
	projects := make([]dto.AllProjectOutput, 0, 10)
	for _, val := range projectsCTX.([]interface{}) {
		project := dto.NewAllProjectOutput(
			fmt.Sprint(val.(map[string]interface{})["name"]),
			fmt.Sprint(val.(map[string]interface{})["code"]),
		)
		projects = append(projects, *project)
	}

	return projects
}
