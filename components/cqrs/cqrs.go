package cqrs

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type FacadeConfig struct {
	GenerateCommandsTopic func(commandName string) string

	CommandHandlers func(commandBus *CommandBus, eventBus *EventBus) []CommandHandler

	CommandsPublisher message.Publisher

	CommandsSubscriberConstructor CommandsSubscriberConstructor

	GenerateEventsTopic func(eventName string) string

	EventHandlers func(commandBus *CommandBus, eventBus *EventBus) []EventHandler

	EventsPublisher message.Publisher

	EventsSubscriberConstructor EventsSubscriberConstructor

	Router *message.Router

	CommandEventMarshaler CommandEventMarshaler

	Logger watermill.LoggerAdapter
}

func (c FacadeConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (c FacadeConfig) EventsEnabled() bool { _ = "STUB: not implemented"; return false }

func (c FacadeConfig) CommandsEnabled() bool { _ = "STUB: not implemented"; return false }

type Facade struct {
	commandsTopic func(commandName string) string
	commandBus    *CommandBus

	eventsTopic func(eventName string) string
	eventBus    *EventBus

	commandEventMarshaler CommandEventMarshaler
}

func (f Facade) CommandBus() *CommandBus { _ = "STUB: not implemented"; return nil }

func (f Facade) EventBus() *EventBus { _ = "STUB: not implemented"; return nil }

func (f Facade) CommandEventMarshaler() CommandEventMarshaler {
	_ = "STUB: not implemented"
	return *new(CommandEventMarshaler)
}

func NewFacade(config FacadeConfig) (*Facade, error) { _ = "STUB: not implemented"; return nil, nil }
