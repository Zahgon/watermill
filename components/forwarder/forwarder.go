package forwarder

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

const defaultForwarderTopic = "forwarder_topic"

type Config struct {
	ForwarderTopic string

	Middlewares []message.HandlerMiddleware

	CloseTimeout time.Duration

	AckWhenCannotUnwrap bool

	Router *message.Router

	Marshaler Marshaler
}

func (c *Config) setDefaults() { _ = "STUB: not implemented"; return }

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

type Forwarder struct {
	router    *message.Router
	publisher message.Publisher
	logger    watermill.LoggerAdapter
	config    Config
}

func NewForwarder(
	subscriberIn message.Subscriber,
	publisherOut message.Publisher,
	logger watermill.LoggerAdapter,
	config Config,
) (*Forwarder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Forwarder) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (f *Forwarder) Close() error { _ = "STUB: not implemented"; return nil }

func (f *Forwarder) Running() chan struct{} { _ = "STUB: not implemented"; return nil }

func (f *Forwarder) forwardMessage(msg *message.Message) error {
	_ = "STUB: not implemented"
	return nil
}
