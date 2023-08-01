package identity

import "github.com/google/uuid"

type IdentityGenerator interface {
	GenerateUUID() uuid.UUID
}

type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) GenerateUUID() uuid.UUID {
	return uuid.New()
}
