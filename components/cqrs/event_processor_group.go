package cqrs

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type EventGroupProcessorConfig struct {
	GenerateSubscribeTopic EventGroupProcessorGenerateSubscribeTopicFn

	SubscriberConstructor EventGroupProcessorSubscriberConstructorFn

	OnHandle EventGroupProcessorOnHandleFn

	AckOnUnknownEvent bool

	Marshaler CommandEventMarshaler

	Logger watermill.LoggerAdapter
}

func (c *EventGroupProcessorConfig) setDefaults() { _ = "STUB: not implemented"; return }

func (c EventGroupProcessorConfig) Validate() error { _ = "STUB: not implemented"; return nil }

type EventGroupProcessorGenerateSubscribeTopicFn func(EventGroupProcessorGenerateSubscribeTopicParams) (string, error)

type EventGroupProcessorGenerateSubscribeTopicParams struct {
	EventGroupName     string
	EventGroupHandlers []GroupEventHandler
}

type EventGroupProcessorSubscriberConstructorFn func(EventGroupProcessorSubscriberConstructorParams) (message.Subscriber, error)

type EventGroupProcessorSubscriberConstructorParams struct {
	EventGroupName     string
	EventGroupHandlers []GroupEventHandler
}

type EventGroupProcessorOnHandleFn func(params EventGroupProcessorOnHandleParams) error

type EventGroupProcessorOnHandleParams struct {
	GroupName string
	Handler   GroupEventHandler

	Event     any
	EventName string

	Message *message.Message
}

type EventGroupProcessor struct {
	router *message.Router

	groupEventHandlers map[string][]GroupEventHandler

	config EventGroupProcessorConfig
}

func NewEventGroupProcessorWithConfig(router *message.Router, config EventGroupProcessorConfig) (*EventGroupProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *EventGroupProcessor) AddHandlersGroup(groupName string, handlers ...GroupEventHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (p EventGroupProcessor) addHandlerToRouter(r *message.Router, groupName string, handlersGroup []GroupEventHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (p EventGroupProcessor) routerHandlerGroupFunc(handlers []GroupEventHandler, groupName string, logger watermill.LoggerAdapter) (message.NoPublishHandlerFunc, error) {
	_ = "STUB: not implemented"
	return *new(message.NoPublishHandlerFunc), nil
}
