package credential

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

type stubDb func(ctx context.Context, receiverID, clientID, organizationID, ssID string) error

func (f stubDb) WriteCredentials(ctx context.Context, receiverID, clientID, organizationID, ssID string) error {
	return f(ctx, receiverID, clientID, organizationID, ssID)
}

func Test_NewCredential(t *testing.T) {
	db := stubDb(func(_ context.Context, _, _, _, _ string) error {
		return nil
	})

	want := Service{
		database: &db,
	}

	got := NewService(&db)

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("got != want: %+v != %+v", got, want)
	}
}

func Test_Save(t *testing.T) {
	db := stubDb(func(_ context.Context, _, _, _, _ string) error {
		return nil
	})
	s := NewService(&db)

	receiverID := "abc"
	c := Credentials{
		clientID:       "123",
		organizationID: "456",
		ssID:           "xyv",
	}

	err := s.Save(context.Background(), receiverID, c)
	if err != nil {
		t.Errorf("s.Save expected nil, got error: %s", err)
	}
}

func Test_Save_Error(t *testing.T) {
	db := stubDb(func(_ context.Context, _, _, _, _ string) error {
		return errors.New("error")
	})
	s := NewService(&db)

	receiverID := "abc"
	c := Credentials{
		clientID:       "123",
		organizationID: "456",
		ssID:           "xyv",
	}

	err := s.Save(context.Background(), receiverID, c)
	if err == nil {
		t.Error("s.Save expected error, got nil")
	}
}
