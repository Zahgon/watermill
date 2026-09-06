package main

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-sqlite/wmsqlitezombiezen"
	"github.com/ThreeDotsLabs/watermill/message"
	_ "modernc.org/sqlite"
	"zombiezen.com/go/sqlite"
)

func main() {
	logger := watermill.NewStdLogger(false, false)

	connectionDSN := "file:ephemeral?mode=memory&cache=shared"
	conn, err := sqlite.OpenConn(connectionDSN)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	publisher, err := wmsqlitezombiezen.NewPublisher(conn, wmsqlitezombiezen.PublisherOptions{
		InitializeSchema: true,
		Logger:           logger,
	})
	if err != nil {
		panic(err)
	}

	subscriber, err := wmsqlitezombiezen.NewSubscriber(connectionDSN, wmsqlitezombiezen.SubscriberOptions{
		InitializeSchema: true,
		Logger:           logger,
	})
	if err != nil {
		panic(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), "example_topic")
	if err != nil {
		panic(err)
	}

	go process(messages)
	publishMessages(publisher)
}

func publishMessages(publisher message.Publisher) { _ = "STUB: not implemented"; return }

func process(messages <-chan *message.Message) { _ = "STUB: not implemented"; return }
