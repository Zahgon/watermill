package message

import (
	"context"
	"sync"
)

func MessageTransformSubscriberDecorator(transform func(*Message)) SubscriberDecorator {
	_ = "STUB: not implemented"
	return *new(SubscriberDecorator)
}

func MessageTransformPublisherDecorator(transform func(*Message)) PublisherDecorator {
	_ = "STUB: not implemented"
	return *new(PublisherDecorator)
}

type messageTransformSubscriberDecorator struct {
	sub Subscriber

	transform   func(*Message)
	subscribeWg sync.WaitGroup
}

func (t *messageTransformSubscriberDecorator) Subscribe(ctx context.Context, topic string) (<-chan *Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *messageTransformSubscriberDecorator) Close() error { _ = "STUB: not implemented"; return nil }

type messageTransformPublisherDecorator struct {
	Publisher
	transform func(*Message)
}

func (d messageTransformPublisherDecorator) Publish(topic string, messages ...*Message) error {
	_ = "STUB: not implemented"
	return nil
}
