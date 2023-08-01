package v1

import (
	"github.com/gin-gonic/gin"
	"handbooks_backend/pkg/common/errors"
	"net/http"
)

var (
	errProjectAll = errors.New("error view projects context")
)

func (h *Handler) allProjects(ctx *gin.Context) {
	projectsCTX, ok := ctx.Get("projects")
	if !ok {
		errorResponse(ctx, http.StatusBadRequest, errProjectAll.Error())
		return
	}
	resp := h.policy.AllProject(ctx, projectsCTX)

	ctx.JSON(http.StatusOK, resp)
}
