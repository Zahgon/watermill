package requeuer

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

const RetriesKey = "_watermill_requeuer_retries"

type Requeuer struct {
	config Config
}

type GeneratePublishTopicParams struct {
	Message *message.Message
}

type Config struct {
	Subscriber message.Subscriber

	SubscribeTopic string

	Publisher message.Publisher

	GeneratePublishTopic func(params GeneratePublishTopicParams) (string, error)

	Delay time.Duration

	Router *message.Router
}

func (c *Config) setDefaults(logger watermill.LoggerAdapter) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) validate() error { _ = "STUB: not implemented"; return nil }

func NewRequeuer(
	config Config,
	logger watermill.LoggerAdapter,
) (*Requeuer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Requeuer) handler(msg *message.Message) error { _ = "STUB: not implemented"; return nil }

func (r *Requeuer) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
