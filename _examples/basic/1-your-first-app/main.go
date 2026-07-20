package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
)

var (
	brokers      = []string{"kafka:9092"}
	consumeTopic = "events"
	publishTopic = "events-processed"

	logger = watermill.NewStdLogger(
		true,
		false,
	)
	marshaler = kafka.DefaultMarshaler{}
)

type event struct {
	ID int `json:"id"`
}

type processedEvent struct {
	ProcessedID int       `json:"processed_id"`
	Time        time.Time `json:"time"`
}

func main() {
	publisher := createPublisher()

	subscriber := createSubscriber("handler_1")

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.Recoverer)

	router.AddHandler(
		"handler_1",
		consumeTopic,
		subscriber,
		publishTopic,
		publisher,
		func(msg *message.Message) ([]*message.Message, error) {
			consumedPayload := event{}
			err := json.Unmarshal(msg.Payload, &consumedPayload)
			if err != nil {

				return nil, err
			}

			fmt.Printf("received event %+v\n", consumedPayload)

			newPayload, err := json.Marshal(processedEvent{
				ProcessedID: consumedPayload.ID,
				Time:        time.Now(),
			})
			if err != nil {
				return nil, err
			}

			newMessage := message.NewMessage(watermill.NewUUID(), newPayload)

			return []*message.Message{newMessage}, nil
		},
	)

	go simulateEvents(publisher)

	if err := router.Run(context.Background()); err != nil {
		panic(err)
	}
}

func createPublisher() message.Publisher { _ = "STUB: not implemented"; return *new(message.Publisher) }

func createSubscriber(consumerGroup string) message.Subscriber {
	_ = "STUB: not implemented"
	return *new(message.Subscriber)
}

func simulateEvents(publisher message.Publisher) { _ = "STUB: not implemented"; return }
