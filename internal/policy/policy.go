package policy

import (
	"handbooks_backend/internal/service"
	"handbooks_backend/pkg/common/core/identity"
)

type Policy struct {
	service  *service.Service
	identity identity.IdentityGenerator
}

func NewPolicy(service *service.Service, identity identity.IdentityGenerator) *Policy {
	return &Policy{
		service:  service,
		identity: identity,
	}
}
