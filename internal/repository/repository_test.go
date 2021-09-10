package repository

import (
	"reflect"
	"testing"

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
