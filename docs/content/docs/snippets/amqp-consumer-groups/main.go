package main

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

var amqpURI = "amqp://guest:guest@rabbitmq:5672/"

func createSubscriber(queueSuffix string) *amqp.Subscriber { _ = "STUB: not implemented"; return nil }

func main() {
	subscriber1 := createSubscriber("test_consumer_group_1")
	messages1, err := subscriber1.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}
	go process("subscriber_1", messages1)

	subscriber2 := createSubscriber("test_consumer_group_2")
	messages2, err := subscriber2.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}

	go process("subscriber_2", messages2)

	publisher, err := amqp.NewPublisher(
		amqp.NewDurablePubSubConfig(
			amqpURI,
			nil,
		),
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	publishMessages(publisher)
}

func publishMessages(publisher message.Publisher) { _ = "STUB: not implemented"; return }

func process(subscriber string, messages <-chan *message.Message) {
	_ = "STUB: not implemented"
	return
}
