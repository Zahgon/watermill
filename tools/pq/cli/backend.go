package cli

import (
	"context"
)

type BackendConfig struct {
	Topic    string
	RawTopic string
}

func (c BackendConfig) Validate() error { _ = "STUB: not implemented"; return nil }

type BackendConstructor func(ctx context.Context, cfg BackendConfig) (Backend, error)

type Backend interface {
	AllMessages(ctx context.Context) ([]Message, error)
	Requeue(ctx context.Context, msg Message) error
	Ack(ctx context.Context, msg Message) error
}
