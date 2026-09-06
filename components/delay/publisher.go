package delay

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type DefaultDelayGeneratorParams struct {
	Topic   string
	Message *message.Message
}

type PublisherConfig struct {
	DefaultDelayGenerator func(params DefaultDelayGeneratorParams) (Delay, error)

	AllowNoDelay bool
}

func NewPublisher(pub message.Publisher, config PublisherConfig) (message.Publisher, error) {
	_ = "STUB: not implemented"
	return *new(message.Publisher), nil
}

type publisher struct {
	pub    message.Publisher
	config PublisherConfig
}

func (p *publisher) Publish(topic string, messages ...*message.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *publisher) Close() error { _ = "STUB: not implemented"; return nil }

func (p *publisher) applyDelay(topic string, msg *message.Message) error {
	_ = "STUB: not implemented"
	return nil
}
