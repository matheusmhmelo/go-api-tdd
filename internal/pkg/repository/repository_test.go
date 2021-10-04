package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"testing"

	"github.com/pashagolub/pgxmock"
)

func TestDatabase_WriteCredentials(t *testing.T) {
	mock, _ := pgxmock.NewConn()
	db := NewDatabase(mock)

	id := uuid.New()
	receiverID := "abc"
	clientID := "123"
	organizationID := "456"
	ssID := "xyv"

	mock.ExpectExec("INSERT INTO credentials").
		WithArgs(id, receiverID, clientID, organizationID, ssID).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err := db.WriteCredentials(context.Background(), id, receiverID, clientID, organizationID, ssID)
	if err != nil {
		t.Errorf("db.WriteCredentials got an error, expected nil: %s", err)
	}
}

func TestDatabase_WriteCredentials_Errors(t *testing.T) {
	t.Run("Unknown SQL error", func(t *testing.T) {
		mock, _ := pgxmock.NewConn()
		db := NewDatabase(mock)

		id := uuid.New()
		receiverID := "abc"
		clientID := "123"
		organizationID := "456"
		ssID := "xyv"

		mock.ExpectExec("INSERT INTO credentials").
			WithArgs(id, receiverID, clientID, organizationID, ssID).
			WillReturnError(errors.New("unknown error"))

		err := db.WriteCredentials(context.Background(), id, receiverID, clientID, organizationID, ssID)
		if err == nil {
			t.Error("db.WriteCredentials expected error, got nil")
		}
	})
}
