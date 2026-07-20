package backend

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/ThreeDotsLabs/watermill/tools/pq/cli"
)

type PostgresMessage struct {
	Offset   int    `db:"offset"`
	UUID     string `db:"uuid"`
	Payload  string `db:"payload"`
	Metadata string `db:"metadata"`
}

type PostgresBackend struct {
	db     *sqlx.DB
	config cli.BackendConfig
}

func NewPostgresBackend(ctx context.Context, config cli.BackendConfig) (*PostgresBackend, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PostgresBackend) AllMessages(ctx context.Context) ([]cli.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PostgresBackend) Requeue(ctx context.Context, msg cli.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostgresBackend) Ack(ctx context.Context, msg cli.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostgresBackend) topic() string { _ = "STUB: not implemented"; return "" }
