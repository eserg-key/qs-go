package v1

import (
	"github.com/gin-gonic/gin"
	"handbooks_backend/internal/policy"
)

const _version = "/api"

type Handler struct {
	policy *policy.Policy
}

func NewHandler(policy *policy.Policy) *Handler {
	return &Handler{policy: policy}
}

func (h *Handler) Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()
	handler.Use(CORS())
	handler.Use(gin.Recovery())

	// Projects
	handler.GET(_version+"/projects", h.userIdentity, h.allProjects)

	// HandBook
	handler.POST(_version+"/handbook", h.userIdentity, h.createHandbook)
	handler.GET(_version+"/handbook", h.userIdentity, h.getHandbooks)
	handler.PATCH(_version+"/handbook/:id", h.userIdentity, h.updateHandbook)
	handler.GET(_version+"/handbook/:id", h.userIdentity, h.getHandbook)
	handler.POST(_version+"/handbook/:id", h.userIdentity, h.updateRowHandbook)
	handler.POST(_version+"/elk", h.testELK)

	return handler
}
