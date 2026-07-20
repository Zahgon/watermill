package main

import (
	"context"
	"database/sql"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-sqlite/wmsqlitemodernc"
	"github.com/ThreeDotsLabs/watermill/message"
	_ "modernc.org/sqlite"
)

func main() {
	db := createDB()
	defer db.Close()
	logger := watermill.NewStdLogger(false, false)

	subscriber, err := wmsqlitemodernc.NewSubscriber(
		db,
		wmsqlitemodernc.SubscriberOptions{
			InitializeSchema: true,
			Logger:           logger,
		},
	)
	if err != nil {
		panic(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), "example_topic")
	if err != nil {
		panic(err)
	}

	go process(messages)

	publisher, err := wmsqlitemodernc.NewPublisher(
		db,
		wmsqlitemodernc.PublisherOptions{
			InitializeSchema: true,
			Logger:           logger,
		},
	)
	if err != nil {
		panic(err)
	}
	publishMessages(publisher)
}

func createDB() *sql.DB { _ = "STUB: not implemented"; return nil }

func publishMessages(publisher message.Publisher) { _ = "STUB: not implemented"; return }

func process(messages <-chan *message.Message) { _ = "STUB: not implemented"; return }
