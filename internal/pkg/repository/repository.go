package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

const insertCredentialsSQL = "INSERT INTO credentials(id, receiver_id, organization_id, software_statement_id, client_id) VALUES ($1, $2, $3, $4, $5)"

type connector interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
}

type Database struct {
	conn connector
}

func NewDatabase(conn connector) *Database {
	return &Database{
		conn: conn,
	}
}

func (d *Database) WriteCredentials(ctx context.Context, id uuid.UUID, receiverID, clientID, organizationID, ssID string) error {
	_, err := d.conn.Exec(ctx, insertCredentialsSQL, id, receiverID, clientID, organizationID, ssID)
	if err != nil {
		return err
	}
	return nil
}
