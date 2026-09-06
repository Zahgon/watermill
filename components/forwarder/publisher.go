package forwarder

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type PublisherConfig struct {
	ForwarderTopic string

	Marshaler Marshaler
}

func (c *PublisherConfig) setDefaults() { _ = "STUB: not implemented"; return }

func (c *PublisherConfig) Validate() error { _ = "STUB: not implemented"; return nil }

type Publisher struct {
	wrappedPublisher message.Publisher
	config           PublisherConfig
}

func NewPublisher(publisher message.Publisher, config PublisherConfig) *Publisher {
	_ = "STUB: not implemented"
	return nil
}

func (p *Publisher) Publish(topic string, messages ...*message.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Publisher) Close() error { _ = "STUB: not implemented"; return nil }
