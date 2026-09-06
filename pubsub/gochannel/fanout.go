package gochannel

import (
	"context"
	"sync"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type FanOut struct {
	internalPubSub *GoChannel
	internalRouter *message.Router

	subscriber message.Subscriber

	logger watermill.LoggerAdapter

	subscribedTopics map[string]struct{}
	subscribedLock   sync.Mutex
}

func NewFanOut(
	subscriber message.Subscriber,
	logger watermill.LoggerAdapter,
) (*FanOut, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FanOut) AddSubscription(topic string) { _ = "STUB: not implemented"; return }

func (f *FanOut) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (f *FanOut) Running() chan struct{} { _ = "STUB: not implemented"; return nil }

func (f *FanOut) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (f *FanOut) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FanOut) Close() error { _ = "STUB: not implemented"; return nil }
