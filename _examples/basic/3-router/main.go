package main

import (
	"context"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

var (
	logger = watermill.NewStdLogger(false, false)
)

func main() {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddPlugin(plugin.SignalsHandler)

	router.AddMiddleware(

		middleware.CorrelationID,

		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: time.Millisecond * 100,
			Logger:          logger,
		}.Middleware,

		middleware.Recoverer,
	)

	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)

	go publishMessages(pubSub)

	handler := router.AddHandler(
		"struct_handler",
		"incoming_messages_topic",
		pubSub,
		"outgoing_messages_topic",
		pubSub,
		structHandler{}.Handler,
	)

	handler.AddMiddleware(func(h message.HandlerFunc) message.HandlerFunc {
		return func(message *message.Message) ([]*message.Message, error) {
			log.Println("executing handler specific middleware for ", message.UUID)

			return h(message)
		}
	})

	router.AddConsumerHandler(
		"print_incoming_messages",
		"incoming_messages_topic",
		pubSub,
		printMessages,
	)

	router.AddConsumerHandler(
		"print_outgoing_messages",
		"outgoing_messages_topic",
		pubSub,
		printMessages,
	)

	ctx := context.Background()
	if err := router.Run(ctx); err != nil {
		panic(err)
	}
}

func publishMessages(publisher message.Publisher) { _ = "STUB: not implemented"; return }

func printMessages(msg *message.Message) error { _ = "STUB: not implemented"; return nil }

type structHandler struct {
}

func (s structHandler) Handler(msg *message.Message) ([]*message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
