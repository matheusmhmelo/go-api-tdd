package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/matheusmhmelo/go-api-tdd/internal/pkg/credential"
)

const partnerHeader = "X-Partner-Id"

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

func (c *Credential) Create(w http.ResponseWriter, r *http.Request) {
	partnerID := r.Header.Get(partnerHeader)
	if partnerID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var creds credential.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = creds.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.service.Save(r.Context(), partnerID, creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
