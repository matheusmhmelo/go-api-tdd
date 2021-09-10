package handler

import (
	"context"

	"github.com/matheusmhmelo/go-api-tdd/internal/pkg/credential"
)

type service interface {
	Save(ctx context.Context, receiverID string, c credential.Credentials) error
}

type Credential struct {
	service service
}

func NewCredential(s service) *Credential {
	return &Credential{
		service: s,
	}
}
