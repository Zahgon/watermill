package gochannel

import (
	"context"
	"sync"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type Config struct {
	OutputChannelBuffer int64

	Persistent bool

	BlockPublishUntilSubscriberAck bool

	PreserveContext bool
}

type GoChannel struct {
	config Config
	logger watermill.LoggerAdapter

	subscribersWg          sync.WaitGroup
	subscribers            map[string][]*subscriber
	subscribersLock        sync.RWMutex
	subscribersByTopicLock sync.Map

	closed     bool
	closedLock sync.Mutex
	closing    chan struct{}

	persistedMessages     map[string][]*message.Message
	persistedMessagesLock sync.RWMutex
}

func NewGoChannel(config Config, logger watermill.LoggerAdapter) *GoChannel {
	_ = "STUB: not implemented"
	return nil
}

func (g *GoChannel) Publish(topic string, messages ...*message.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *GoChannel) waitForAckFromSubscribers(msg *message.Message, ackedByConsumer <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (g *GoChannel) sendMessage(topic string, message *message.Message) (<-chan struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GoChannel) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *GoChannel) addSubscriber(topic string, s *subscriber) { _ = "STUB: not implemented"; return }

func (g *GoChannel) removeSubscriber(topic string, toRemove *subscriber) {
	_ = "STUB: not implemented"
	return
}

func (g *GoChannel) topicSubscribers(topic string) []*subscriber {
	_ = "STUB: not implemented"
	return nil
}

func (g *GoChannel) isClosed() bool { _ = "STUB: not implemented"; return false }

func (g *GoChannel) Close() error { _ = "STUB: not implemented"; return nil }

type subscriber struct {
	ctx context.Context

	uuid string

	sending       sync.Mutex
	outputChannel chan *message.Message

	logger  watermill.LoggerAdapter
	closed  bool
	closing chan struct{}

	preserveContext bool
}

func (s *subscriber) Close() { _ = "STUB: not implemented"; return }

func (s *subscriber) sendMessageToSubscriber(msg *message.Message, logFields watermill.LogFields) {
	_ = "STUB: not implemented"
	return
}
