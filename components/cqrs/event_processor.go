package cqrs

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type EventProcessorConfig struct {
	GenerateSubscribeTopic EventProcessorGenerateSubscribeTopicFn

	SubscriberConstructor EventProcessorSubscriberConstructorFn

	OnHandle EventProcessorOnHandleFn

	AckOnUnknownEvent bool

	Marshaler CommandEventMarshaler

	Logger watermill.LoggerAdapter

	disableRouterAutoAddHandlers bool
}

func (c *EventProcessorConfig) setDefaults() { _ = "STUB: not implemented"; return }

func (c EventProcessorConfig) Validate() error { _ = "STUB: not implemented"; return nil }

type EventProcessorGenerateSubscribeTopicFn func(EventProcessorGenerateSubscribeTopicParams) (string, error)

type EventProcessorGenerateSubscribeTopicParams struct {
	EventName    string
	EventHandler EventHandler
}

type EventProcessorSubscriberConstructorFn func(EventProcessorSubscriberConstructorParams) (message.Subscriber, error)

type EventProcessorSubscriberConstructorParams struct {
	EventName    string
	HandlerName  string
	EventHandler EventHandler
}

type EventProcessorOnHandleFn func(params EventProcessorOnHandleParams) error

type EventProcessorOnHandleParams struct {
	Handler EventHandler

	Event     any
	EventName string

	Message *message.Message
}

type EventProcessor struct {
	router   *message.Router
	handlers []EventHandler
	config   EventProcessorConfig
}

func NewEventProcessorWithConfig(router *message.Router, config EventProcessorConfig) (*EventProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewEventProcessor(
	individualHandlers []EventHandler,
	generateTopic func(eventName string) string,
	subscriberConstructor EventsSubscriberConstructor,
	marshaler CommandEventMarshaler,
	logger watermill.LoggerAdapter,
) (*EventProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type EventsSubscriberConstructor func(handlerName string) (message.Subscriber, error)

func (p *EventProcessor) AddHandlers(handlers ...EventHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *EventProcessor) AddHandler(handler EventHandler) (*message.Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p EventProcessor) AddHandlersToRouter(r *message.Router) error {
	_ = "STUB: not implemented"
	return nil
}

func (p EventProcessor) addHandlerToRouter(r *message.Router, handler EventHandler) (*message.Handler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p EventProcessor) Handlers() []EventHandler { _ = "STUB: not implemented"; return nil }

func addHandlerToRouter(logger watermill.LoggerAdapter, r *message.Router, handlerName string, topicName string, handlerFunc message.NoPublishHandlerFunc, subscriber message.Subscriber) *message.Handler {
	_ = "STUB: not implemented"
	return nil
}

func (p EventProcessor) routerHandlerFunc(handler EventHandler, logger watermill.LoggerAdapter) (message.NoPublishHandlerFunc, error) {
	_ = "STUB: not implemented"
	return *new(message.NoPublishHandlerFunc), nil
}

func validateEvent(event interface{}) error { _ = "STUB: not implemented"; return nil }
