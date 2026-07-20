package main

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	watermill_http "github.com/ThreeDotsLabs/watermill-http/pkg/http"
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
)

var (
	logger = watermill.NewStdLogger(false, false)
)

func filterMessages(acceptedTypes ...string) message.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(message.HandlerFunc)
}

func main() {
	publisher, err := watermill_http.NewPublisher(watermill_http.PublisherConfig{
		MarshalMessageFunc: watermill_http.DefaultMarshalMessageFunc,
	}, logger)
	if err != nil {
		panic(err)
	}

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:     []string{"kafka:9092"},
			Unmarshaler: kafka.DefaultMarshaler{},
		},
		logger,
	)
	if err != nil {
		panic(err)
	}

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	topic := "kafka_to_http_example"
	url := "http://webhooks-server:8001/"

	router.AddHandler("foo", topic, subscriber, url+"foo", publisher, filterMessages("Foo"))
	router.AddHandler("foo_or_bar", topic, subscriber, url+"foo_or_bar", publisher, filterMessages("Foo", "Bar"))
	router.AddHandler("all", topic, subscriber, url+"all", publisher, filterMessages("Foo", "Bar", "Baz"))
	router.AddPlugin(plugin.SignalsHandler)

	err = router.Run(context.Background())
	if err != nil {
		panic(err)
	}
}
