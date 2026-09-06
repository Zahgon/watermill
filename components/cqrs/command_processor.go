package cqrs

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type CommandProcessorConfig struct {
	GenerateSubscribeTopic CommandProcessorGenerateSubscribeTopicFn

	SubscriberConstructor CommandProcessorSubscriberConstructorFn

	OnHandle CommandProcessorOnHandleFn

	Marshaler CommandEventMarshaler

	Logger watermill.LoggerAdapter

	AckCommandHandlingErrors bool

	disableRouterAutoAddHandlers bool
}

func (c *CommandProcessorConfig) setDefaults() { _ = "STUB: not implemented"; return }

func (c CommandProcessorConfig) Validate() error { _ = "STUB: not implemented"; return nil }

type CommandProcessorGenerateSubscribeTopicFn func(CommandProcessorGenerateSubscribeTopicParams) (string, error)

type CommandProcessorGenerateSubscribeTopicParams struct {
	CommandName    string
	CommandHandler CommandHandler
}

type CommandProcessorSubscriberConstructorFn func(CommandProcessorSubscriberConstructorParams) (message.Subscriber, error)

type CommandProcessorSubscriberConstructorParams struct {
	CommandName string
	HandlerName string
	Handler     CommandHandler
}

type CommandProcessorOnHandleFn func(params CommandProcessorOnHandleParams) error

type CommandProcessorOnHandleParams struct {
	Handler CommandHandler

	CommandName string
	Command     any

	Message *message.Message
}

type CommandProcessor struct {
	router *message.Router

	handlers []CommandHandler

	config CommandProcessorConfig
}

func NewCommandProcessorWithConfig(router *message.Router, config CommandProcessorConfig) (*CommandProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewCommandProcessor(
	handlers []CommandHandler,
	generateTopic func(commandName string) string,
	subscriberConstructor CommandsSubscriberConstructor,
	marshaler CommandEventMarshaler,
	logger watermill.LoggerAdapter,
) (*CommandProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CommandsSubscriberConstructor func(handlerName string) (message.Subscriber, error)

func (p *CommandProcessor) AddHandlers(handlers ...CommandHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *CommandProcessor) AddHandler(handler CommandHandler) (*message.Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DuplicateCommandHandlerError struct {
	CommandName string
}

func (d DuplicateCommandHandlerError) Error() string { _ = "STUB: not implemented"; return "" }

func (p CommandProcessor) AddHandlersToRouter(r *message.Router) error {
	_ = "STUB: not implemented"
	return nil
}

func (p CommandProcessor) addHandlerToRouter(r *message.Router, handler CommandHandler) (*message.Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p CommandProcessor) Handlers() []CommandHandler { _ = "STUB: not implemented"; return nil }

func (p CommandProcessor) routerHandlerFunc(handler CommandHandler, logger watermill.LoggerAdapter) (message.NoPublishHandlerFunc, error) {
	_ = "STUB: not implemented"
	return *new(message.NoPublishHandlerFunc), nil
}

func (p CommandProcessor) validateCommand(cmd interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
