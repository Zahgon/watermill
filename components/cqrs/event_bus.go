package cqrs

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type EventBusConfig struct {
	GeneratePublishTopic GenerateEventPublishTopicFn

	OnPublish OnEventSendFn

	Marshaler CommandEventMarshaler

	Logger watermill.LoggerAdapter
}

func (c *EventBusConfig) setDefaults() { _ = "STUB: not implemented"; return }

func (c EventBusConfig) Validate() error { _ = "STUB: not implemented"; return nil }

type GenerateEventPublishTopicFn func(GenerateEventPublishTopicParams) (string, error)

type GenerateEventPublishTopicParams struct {
	EventName string
	Event     any
}

type OnEventSendFn func(params OnEventSendParams) error

type OnEventSendParams struct {
	EventName string
	Event     any

	Message *message.Message
}

type EventBus struct {
	publisher message.Publisher
	config    EventBusConfig
}

func NewEventBus(
	publisher message.Publisher,
	generateTopic func(eventName string) string,
	marshaler CommandEventMarshaler,
) (*EventBus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewEventBusWithConfig(publisher message.Publisher, config EventBusConfig) (*EventBus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c EventBus) Publish(ctx context.Context, event any) error {
	_ = "STUB: not implemented"
	return nil
}
