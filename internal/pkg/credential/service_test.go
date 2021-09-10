package credential

import (
	"context"
	"reflect"
	"testing"
)

type stubDb func(ctx context.Context, receiverID, clientID, organizationID, ssID string) error

func (f stubDb) WriteCredentials(_ context.Context, _, _, _, _ string) error {
	return nil
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
