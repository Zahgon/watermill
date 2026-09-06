package message

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/ThreeDotsLabs/watermill"
)

var (
	ErrOutputInNoPublisherHandler = errors.New("returned output messages in a handler without publisher")
)

type HandlerFunc func(msg *Message) ([]*Message, error)

type NoPublishHandlerFunc func(msg *Message) error

var PassthroughHandler HandlerFunc = func(msg *Message) ([]*Message, error) {
	return []*Message{msg}, nil
}

type HandlerMiddleware func(h HandlerFunc) HandlerFunc

type RouterPlugin func(*Router) error

type PublisherDecorator func(pub Publisher) (Publisher, error)

type SubscriberDecorator func(sub Subscriber) (Subscriber, error)

type RouterConfig struct {
	CloseTimeout time.Duration
}

func (c *RouterConfig) setDefaults() { _ = "STUB: not implemented"; return }

func (c RouterConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func NewRouter(config RouterConfig, logger watermill.LoggerAdapter) (*Router, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDefaultRouter(logger watermill.LoggerAdapter) *Router {
	_ = "STUB: not implemented"
	return nil
}

func newRouter(config RouterConfig, logger watermill.LoggerAdapter) *Router {
	_ = "STUB: not implemented"
	return nil
}

type middleware struct {
	Handler       HandlerMiddleware
	HandlerName   string
	IsRouterLevel bool
}

type Router struct {
	config RouterConfig

	middlewares     []middleware
	middlewaresLock *sync.RWMutex

	plugins []RouterPlugin

	handlers     map[string]*handler
	handlersLock *sync.RWMutex

	handlersWg *sync.WaitGroup

	runningHandlersWg     *sync.WaitGroup
	runningHandlersWgLock *sync.Mutex

	handlerAdded chan struct{}

	closingInProgressCh chan struct{}
	closedCh            chan struct{}
	closed              bool
	closedLock          sync.Mutex

	logger watermill.LoggerAdapter

	publisherDecorators  []PublisherDecorator
	subscriberDecorators []SubscriberDecorator

	isRunning bool
	running   chan struct{}
}

func (r *Router) Logger() watermill.LoggerAdapter {
	_ = "STUB: not implemented"
	return *new(watermill.LoggerAdapter)
}

func (r *Router) AddMiddleware(m ...HandlerMiddleware) { _ = "STUB: not implemented"; return }

func (r *Router) addRouterLevelMiddleware(m ...HandlerMiddleware) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) addHandlerLevelMiddleware(handlerName string, m ...HandlerMiddleware) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) AddPlugin(p ...RouterPlugin) { _ = "STUB: not implemented"; return }

func (r *Router) AddPublisherDecorators(dec ...PublisherDecorator) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) AddSubscriberDecorators(dec ...SubscriberDecorator) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) Handlers() map[string]HandlerFunc { _ = "STUB: not implemented"; return nil }

type DuplicateHandlerNameError struct {
	HandlerName string
}

func (d DuplicateHandlerNameError) Error() string { _ = "STUB: not implemented"; return "" }

func (r *Router) AddHandler(
	handlerName string,
	subscribeTopic string,
	subscriber Subscriber,
	publishTopic string,
	publisher Publisher,
	handlerFunc HandlerFunc,
) *Handler {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) AddConsumerHandler(
	handlerName string,
	subscribeTopic string,
	subscriber Subscriber,
	handlerFunc NoPublishHandlerFunc,
) *Handler {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) AddNoPublisherHandler(
	handlerName string,
	subscribeTopic string,
	subscriber Subscriber,
	handlerFunc NoPublishHandlerFunc,
) *Handler {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) Run(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

func (r *Router) RunHandlers(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *Router) watchAllHandlersStopped(ctx context.Context) { _ = "STUB: not implemented"; return }

func (r *Router) Running() chan struct{} { _ = "STUB: not implemented"; return nil }

func (r *Router) IsRunning() bool { _ = "STUB: not implemented"; return false }

func (r *Router) Close() error { _ = "STUB: not implemented"; return nil }

func (r *Router) waitForHandlers() bool { _ = "STUB: not implemented"; return false }

func (r *Router) IsClosed() bool { _ = "STUB: not implemented"; return false }

type handler struct {
	name   string
	logger watermill.LoggerAdapter

	subscriber     Subscriber
	subscribeTopic string
	subscriberName string

	publisher     Publisher
	publishTopic  string
	publisherName string

	handlerFunc HandlerFunc

	runningHandlersWg     *sync.WaitGroup
	runningHandlersWgLock *sync.Mutex

	messagesCh <-chan *Message

	started   bool
	startedCh chan struct{}

	stopFn         context.CancelFunc
	stopped        chan struct{}
	routersCloseCh chan struct{}
}

func (h *handler) run(ctx context.Context, middlewares []middleware) {
	_ = "STUB: not implemented"
	return
}

type Handler struct {
	router  *Router
	handler *handler
}

func (h *Handler) AddMiddleware(m ...HandlerMiddleware) { _ = "STUB: not implemented"; return }

func (h *Handler) Started() chan struct{} { _ = "STUB: not implemented"; return nil }

func (h *Handler) Stop() { _ = "STUB: not implemented"; return }

func (h *Handler) Stopped() chan struct{} { _ = "STUB: not implemented"; return nil }

func (r *Router) decorateHandlerPublisher(h *handler) error { _ = "STUB: not implemented"; return nil }

func (r *Router) decorateHandlerSubscriber(h *handler) error { _ = "STUB: not implemented"; return nil }

func (h *handler) addHandlerContext(messages ...*Message) { _ = "STUB: not implemented"; return }

func (h *handler) handleClose(ctx context.Context) { _ = "STUB: not implemented"; return }

func (h *handler) handleMessage(msg *Message, handler HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) publishProducedMessages(producedMessages Messages, msgFields watermill.LogFields) error {
	_ = "STUB: not implemented"
	return nil
}

type disabledPublisher struct{}

func (disabledPublisher) Publish(topic string, messages ...*Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (disabledPublisher) Close() error { _ = "STUB: not implemented"; return nil }
