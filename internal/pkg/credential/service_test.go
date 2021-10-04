package credential

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"testing"
)

type stubDb struct {
	funcWriteCredentials func(context.Context, uuid.UUID, string, string, string, string) error
}

func (f *stubDb) WriteCredentials(ctx context.Context, id uuid.UUID, receiverID, clientID, organizationID, ssID string) error {
	return f.funcWriteCredentials(ctx, id, receiverID, clientID, organizationID, ssID)
}

func Test_Save(t *testing.T) {
	ctx := context.Background()
	receiverID := "abc"
	cred := Credentials{
		ClientID:            "123",
		OrganizationID:      "456",
		SoftwareStatementID: "xyv",
	}

	id := uuid.New()
	fakeUuid := func() uuid.UUID {
		return id
	}

	db := stubDb{
		funcWriteCredentials: func(_ context.Context, rId uuid.UUID, rRID, rCID, rOID, rSSID string) error {
			if id != rId ||
				receiverID != rRID ||
				cred.ClientID != rCID ||
				cred.OrganizationID != rOID ||
				cred.SoftwareStatementID != rSSID {
				t.Fatal("Database received unexpected values")
			}
			return nil
		},
	}

	s := NewService(&db, fakeUuid)

	err := s.Save(ctx, receiverID, cred)
	if err != nil {
		t.Errorf("s.Save expected nil, got error: %s", err)
	}
}

func Test_Save_Error(t *testing.T) {
	db := stubDb{
		funcWriteCredentials: func(_ context.Context, _ uuid.UUID, _, _, _, _ string) error {
			return errors.New("error")
		},
	}
	s := NewService(&db, uuid.New)

	ctx := context.Background()
	receiverID := "abc"
	cred := Credentials{
		ClientID:            "123",
		OrganizationID:      "456",
		SoftwareStatementID: "xyv",
	}

	err := s.Save(ctx, receiverID, cred)
	if err == nil {
		t.Error("s.Save expected error, got nil")
	}
}
