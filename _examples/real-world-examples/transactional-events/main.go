package main

import (
	"context"
	stdSQL "database/sql"
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
)

var (
	logger     = watermill.NewStdLogger(false, false)
	kafkaTopic = "events"
	mysqlTable = "events"
)

func main() {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.Recoverer)

	db := createDB()

	subscriber := createSubscriber(db)
	publisher := createPublisher()

	router.AddHandler(
		"mysql-to-kafka",
		mysqlTable,
		subscriber,
		kafkaTopic,
		publisher,
		func(msg *message.Message) ([]*message.Message, error) {
			consumedEvent := event{}
			err := json.Unmarshal(msg.Payload, &consumedEvent)
			if err != nil {
				return nil, err
			}

			log.Printf("received event %+v with UUID %s", consumedEvent, msg.UUID)

			return []*message.Message{msg}, nil
		},
	)

	go func() {
		<-router.Running()
		simulateEvents(db)
	}()

	if err := router.Run(context.Background()); err != nil {
		panic(err)
	}
}

func createDB() *stdSQL.DB { _ = "STUB: not implemented"; return nil }

func createSubscriber(db *stdSQL.DB) message.Subscriber {
	_ = "STUB: not implemented"
	return *new(message.Subscriber)
}

func createPublisher() message.Publisher { _ = "STUB: not implemented"; return *new(message.Publisher) }

type event struct {
	Name       string `json:"name"`
	OccurredAt string `json:"occurred_at"`
}

func simulateEvents(db *stdSQL.DB) { _ = "STUB: not implemented"; return }

func publishEvent(tx *stdSQL.Tx) error { _ = "STUB: not implemented"; return nil }
