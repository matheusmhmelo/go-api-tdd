package credential

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"testing"
)

type stubDb func(ctx context.Context, id uuid.UUID, receiverID, clientID, organizationID, ssID string) error

func (f stubDb) WriteCredentials(ctx context.Context, id uuid.UUID, receiverID, clientID, organizationID, ssID string) error {
	return f(ctx, id, receiverID, clientID, organizationID, ssID)
}

func Test_Save(t *testing.T) {
	db := stubDb(func(_ context.Context, _ uuid.UUID, _, _, _, _ string) error {
		return nil
	})
	s := NewService(&db)

	receiverID := "abc"
	c := Credentials{
		ClientID:            "123",
		OrganizationID:      "456",
		SoftwareStatementID: "xyv",
	}

	err := s.Save(context.Background(), receiverID, c)
	if err != nil {
		t.Errorf("s.Save expected nil, got error: %s", err)
	}
}

func Test_Save_Error(t *testing.T) {
	db := stubDb(func(_ context.Context, _ uuid.UUID, _, _, _, _ string) error {
		return errors.New("error")
	})
	s := NewService(&db)

	receiverID := "abc"
	c := Credentials{
		ClientID:            "123",
		OrganizationID:      "456",
		SoftwareStatementID: "xyv",
	}

	err := s.Save(context.Background(), receiverID, c)
	if err == nil {
		t.Error("s.Save expected error, got nil")
	}
}
