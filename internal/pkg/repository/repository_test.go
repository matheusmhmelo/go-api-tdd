package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/pashagolub/pgxmock"
)

func Test_NewDatabase(t *testing.T) {
	mock, _ := pgxmock.NewConn()

	got := NewDatabase(mock)
	want := Database{
		conn: mock,
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("got != want: %+v != %+v", got, want)
	}
}

func TestDatabase_WriteCredentials(t *testing.T) {
	mock, _ := pgxmock.NewConn()
	db := NewDatabase(mock)

	receiverID := "abc"
	clientID := "123"
	organizationID := "456"
	ssID := "xyv"

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO credentials").
		WithArgs(receiverID, clientID, organizationID, ssID).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))
	mock.ExpectCommit()
	mock.ExpectRollback()

	err := db.WriteCredentials(context.Background(), receiverID, clientID, organizationID, ssID)
	if err != nil {
		t.Errorf("db.WriteCredentials got an error, expected nil: %s", err)
	}
}

func TestDatabase_WriteCredentials_Errors(t *testing.T) {
	t.Run("Unknown SQL error", func(t *testing.T) {
		mock, _ := pgxmock.NewConn()
		db := NewDatabase(mock)

		receiverID := "abc"
		clientID := "123"
		organizationID := "456"
		ssID := "xyv"

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO credentials").
			WithArgs(receiverID, clientID, organizationID, ssID).
			WillReturnError(errors.New("unknown error"))
		mock.ExpectCommit()
		mock.ExpectRollback()

		err := db.WriteCredentials(context.Background(), receiverID, clientID, organizationID, ssID)
		if err == nil {
			t.Error("db.WriteCredentials expected error, got nil")
		}
	})
	t.Run("Credentials Already Exists", func(t *testing.T) {
		mock, _ := pgxmock.NewConn()
		db := NewDatabase(mock)

		receiverID := "abc"
		clientID := "123"
		organizationID := "456"
		ssID := "xyv"

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO credentials").
			WithArgs(receiverID, clientID, organizationID, ssID).
			WillReturnError(&pgconn.PgError{Code: pgerrcode.UniqueViolation})
		mock.ExpectCommit()
		mock.ExpectRollback()

		err := db.WriteCredentials(context.Background(), receiverID, clientID, organizationID, ssID)
		if err == nil {
			t.Error("db.WriteCredentials expected error, got nil")
		}
	})
}
