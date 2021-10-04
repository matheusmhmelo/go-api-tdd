package credential

import (
	"context"
	"github.com/google/uuid"
)

type database interface {
	WriteCredentials(ctx context.Context, id uuid.UUID, receiverID, clientID, organizationID, ssID string) error
}

type Service struct {
	database database
	getId    func() uuid.UUID
}

func NewService(db database, id func() uuid.UUID) *Service {
	return &Service{
		database: db,
		getId:    id,
	}
}

func (s *Service) Save(ctx context.Context, receiverID string, c Credentials) error {
	id := s.getId()

	return s.database.WriteCredentials(
		ctx,
		id,
		receiverID,
		c.ClientID,
		c.OrganizationID,
		c.SoftwareStatementID,
	)
}
