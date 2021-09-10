package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
)

const insertCredentialsSQL = "INSERT INTO credentials(receiver_id, organization_id, software_statement_id, client_id) VALUES ($1, $2, $3, $4)"

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

func (d *Database) WriteCredentials(ctx context.Context, receiverID, clientID, organizationID, ssID string) error {
	return d.conn.BeginFunc(ctx, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, insertCredentialsSQL, receiverID, clientID, organizationID, ssID)
		if err != nil {
			return err
		}
		return nil
	})
}
