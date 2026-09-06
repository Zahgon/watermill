package main

import (
	"context"
	stdSQL "database/sql"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-sql/v4/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	db := createDB()
	logger := watermill.NewStdLogger(false, false)

	subscriber, err := sql.NewSubscriber(
		sql.BeginnerFromStdSQL(db),
		sql.SubscriberConfig{
			SchemaAdapter:    sql.DefaultMySQLSchema{},
			OffsetsAdapter:   sql.DefaultMySQLOffsetsAdapter{},
			InitializeSchema: true,
		},
		logger,
	)
	if err != nil {
		panic(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), "example_topic")
	if err != nil {
		panic(err)
	}

	go process(messages)

	publisher, err := sql.NewPublisher(
		sql.BeginnerFromStdSQL(db),
		sql.PublisherConfig{
			SchemaAdapter: sql.DefaultMySQLSchema{},
		},
		logger,
	)
	if err != nil {
		panic(err)
	}

	publishMessages(publisher)
}

func createDB() *stdSQL.DB { _ = "STUB: not implemented"; return nil }

func publishMessages(publisher message.Publisher) { _ = "STUB: not implemented"; return }

func process(messages <-chan *message.Message) { _ = "STUB: not implemented"; return }
