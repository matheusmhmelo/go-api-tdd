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
}

func NewService(db database) *Service {
	return &Service{
		database: db,
	}
}

func (s *Service) Save(ctx context.Context, receiverID string, c Credentials) error {
	var id uuid.UUID

	return s.database.WriteCredentials(
		ctx,
		id,
		receiverID,
		c.ClientID,
		c.OrganizationID,
		c.SoftwareStatementID,
	)
}
