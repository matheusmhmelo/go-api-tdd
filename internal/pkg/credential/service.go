package credential

import "context"

type database interface {
	WriteCredentials(ctx context.Context, receiverID, clientID, organizationID, ssID string) error
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
	return s.database.WriteCredentials(
		ctx,
		receiverID,
		c.ClientID,
		c.OrganizationID,
		c.SoftwareStatementID,
	)
}
