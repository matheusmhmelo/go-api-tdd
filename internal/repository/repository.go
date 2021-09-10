package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type connector interface {
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
}

type Database struct {
	conn connector
}

func NewDatabase(conn connector) *Database {
	return &Database{
		conn: conn,
	}
}
