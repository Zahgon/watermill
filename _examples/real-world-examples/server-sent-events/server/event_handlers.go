package main

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

const (
	PostCreatedTopic = "post-created"
	PostUpdatedTopic = "post-updated"
	FeedUpdatedTopic = "feed-updated"
)

func SetupMessageRouter(
	feedsStorage FeedsStorage,
	logger watermill.LoggerAdapter,
) (message.Publisher, message.Subscriber, error) {
	_ = "STUB: not implemented"
	return *new(message.Publisher), *new(message.Subscriber), nil
}

func createFeedUpdatedEvents(tags []string) ([]*message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Publisher struct {
	publisher message.Publisher
}

func (p Publisher) Publish(topic string, event interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
