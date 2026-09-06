package fanin

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type Config struct {
	SourceTopics []string

	TargetTopic string

	CloseTimeout time.Duration
}

type FanIn struct {
	router *message.Router
	config Config
	logger watermill.LoggerAdapter
}

func (c *Config) setDefaults() { _ = "STUB: not implemented"; return }

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func NewFanIn(
	subscriber message.Subscriber,
	publisher message.Publisher,
	config Config,
	logger watermill.LoggerAdapter,
) (*FanIn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FanIn) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (f *FanIn) Running() chan struct{} { _ = "STUB: not implemented"; return nil }

func (f *FanIn) Close() error { _ = "STUB: not implemented"; return nil }
