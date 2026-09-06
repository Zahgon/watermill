package main

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-googlecloud/v2/pkg/googlecloud"
	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	logger := watermill.NewStdLogger(false, false)
	subscriber, err := googlecloud.NewSubscriber(
		googlecloud.SubscriberConfig{

			GenerateSubscriptionName: func(topic string) string {
				return "test-sub_" + topic
			},
			ProjectID: "test-project",
		},
		logger,
	)
	if err != nil {
		panic(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}

	go process(messages)

	publisher, err := googlecloud.NewPublisher(googlecloud.PublisherConfig{
		ProjectID: "test-project",
	}, logger)
	if err != nil {
		panic(err)
	}

	publishMessages(publisher)
}

func publishMessages(publisher message.Publisher) { _ = "STUB: not implemented"; return }

func process(messages <-chan *message.Message) { _ = "STUB: not implemented"; return }
