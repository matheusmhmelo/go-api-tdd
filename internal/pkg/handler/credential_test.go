package handler

import (
	"context"
	"reflect"
	"testing"

	"github.com/matheusmhmelo/go-api-tdd/internal/pkg/credential"
)

type stubService func(ctx context.Context, receiverID string, c credential.Credentials) error

func (s stubService) Save(ctx context.Context, receiverID string, c credential.Credentials) error {
	return s(ctx, receiverID, c)
}

func Test_NewHandler(t *testing.T) {
	s := stubService(func(_ context.Context, _ string, _ credential.Credentials) error {
		return nil
	})

	want := Credential{
		service: &s,
	}

	got := NewCredential(&s)

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("got != want: %+v != %+v", got, want)
	}
}
