package cqrs

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"

	"github.com/ThreeDotsLabs/watermill/message"
)

type CommandBusConfig struct {
	GeneratePublishTopic CommandBusGeneratePublishTopicFn

	OnSend CommandBusOnSendFn

	Marshaler CommandEventMarshaler

	Logger watermill.LoggerAdapter
}

func (c *CommandBusConfig) setDefaults() { _ = "STUB: not implemented"; return }

func (c CommandBusConfig) Validate() error { _ = "STUB: not implemented"; return nil }

type CommandBusGeneratePublishTopicFn func(CommandBusGeneratePublishTopicParams) (string, error)

type CommandBusGeneratePublishTopicParams struct {
	CommandName string
	Command     any
}

type CommandBusOnSendFn func(params CommandBusOnSendParams) error

type CommandBusOnSendParams struct {
	CommandName string
	Command     any

	Message *message.Message
}

type CommandBus struct {
	publisher message.Publisher

	config CommandBusConfig
}

func NewCommandBusWithConfig(publisher message.Publisher, config CommandBusConfig) (*CommandBus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewCommandBus(
	publisher message.Publisher,
	generateTopic func(commandName string) string,
	marshaler CommandEventMarshaler,
) (*CommandBus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c CommandBus) Send(ctx context.Context, cmd any) error { _ = "STUB: not implemented"; return nil }

func (c CommandBus) SendWithModifiedMessage(ctx context.Context, cmd any, modify func(*message.Message) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CommandBus) newMessage(ctx context.Context, command any) (*message.Message, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}
