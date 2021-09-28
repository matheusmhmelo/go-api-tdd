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
	receiverID := "abc"
	c := Credentials{
		ClientID:            "123",
		OrganizationID:      "456",
		SoftwareStatementID: "xyv",
	}

	id := uuid.New()
	spyUuid := func() uuid.UUID {
		return id
	}

	db := stubDb(func(_ context.Context, rId uuid.UUID, rRID, rCID, rOID, rSSID string) error {
		if id != rId ||
			receiverID != rRID ||
			c.ClientID != rCID ||
			c.OrganizationID != rOID ||
			c.SoftwareStatementID != rSSID {
			t.Fatal("Database received unexpected values")
		}
		return nil
	})

	s := NewService(&db, spyUuid)

	err := s.Save(context.Background(), receiverID, c)
	if err != nil {
		t.Errorf("s.Save expected nil, got error: %s", err)
	}
}

func Test_Save_Error(t *testing.T) {
	db := stubDb(func(_ context.Context, _ uuid.UUID, _, _, _, _ string) error {
		return errors.New("error")
	})
	s := NewService(&db, func() uuid.UUID {
		return uuid.New()
	})

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
