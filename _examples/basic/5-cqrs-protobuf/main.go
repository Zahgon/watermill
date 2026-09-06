package main

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

type BookRoomHandler struct {
	eventBus *cqrs.EventBus
}

func (b BookRoomHandler) Handle(ctx context.Context, cmd *BookRoom) error {
	_ = "STUB: not implemented"
	return nil
}

type OrderBeerOnRoomBooked struct {
	commandBus *cqrs.CommandBus
}

func (o OrderBeerOnRoomBooked) Handle(ctx context.Context, event *RoomBooked) error {
	_ = "STUB: not implemented"
	return nil
}

type OrderBeerHandler struct {
	eventBus *cqrs.EventBus
}

func (o OrderBeerHandler) Handle(ctx context.Context, cmd *OrderBeer) error {
	_ = "STUB: not implemented"
	return nil
}

type BookingsFinancialReport struct {
	handledBookings map[string]struct{}
	totalCharge     int64
	lock            sync.Mutex
}

func NewBookingsFinancialReport() *BookingsFinancialReport { _ = "STUB: not implemented"; return nil }

func (b *BookingsFinancialReport) Handle(ctx context.Context, event *RoomBooked) error {
	_ = "STUB: not implemented"
	return nil
}

var amqpAddress = "amqp://guest:guest@rabbitmq:5672/"

func main() {
	logger := watermill.NewSlogLoggerWithLevelMapping(nil, map[slog.Level]slog.Level{
		slog.LevelInfo: slog.LevelDebug,
	})

	cqrsMarshaler := cqrs.ProtoMarshaler{

		GenerateName: cqrs.StructName,
	}

	generateEventsTopic := func(eventName string) string {
		return "events." + eventName
	}

	generateCommandsTopic := func(commandName string) string {
		return "commands." + commandName
	}

	commandsAMQPConfig := amqp.NewDurableQueueConfig(amqpAddress)
	commandsPublisher, err := amqp.NewPublisher(commandsAMQPConfig, logger)
	if err != nil {
		panic(err)
	}
	commandsSubscriber, err := amqp.NewSubscriber(commandsAMQPConfig, logger)
	if err != nil {
		panic(err)
	}

	eventsPublisher, err := amqp.NewPublisher(amqp.NewDurablePubSubConfig(amqpAddress, nil), logger)
	if err != nil {
		panic(err)
	}

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddMiddleware(middleware.Recoverer)

	commandBus, err := cqrs.NewCommandBusWithConfig(commandsPublisher, cqrs.CommandBusConfig{
		GeneratePublishTopic: func(params cqrs.CommandBusGeneratePublishTopicParams) (string, error) {
			return generateCommandsTopic(params.CommandName), nil
		},
		OnSend: func(params cqrs.CommandBusOnSendParams) error {
			logger.Info("Sending command", watermill.LogFields{
				"command_name": params.CommandName,
			})

			params.Message.Metadata.Set("sent_at", time.Now().String())

			return nil
		},
		Marshaler: cqrsMarshaler,
		Logger:    logger,
	})
	if err != nil {
		panic(err)
	}

	commandProcessor, err := cqrs.NewCommandProcessorWithConfig(
		router,
		cqrs.CommandProcessorConfig{
			GenerateSubscribeTopic: func(params cqrs.CommandProcessorGenerateSubscribeTopicParams) (string, error) {
				return generateCommandsTopic(params.CommandName), nil
			},
			SubscriberConstructor: func(params cqrs.CommandProcessorSubscriberConstructorParams) (message.Subscriber, error) {

				return commandsSubscriber, nil
			},
			OnHandle: func(params cqrs.CommandProcessorOnHandleParams) error {
				start := time.Now()

				err := params.Handler.Handle(params.Message.Context(), params.Command)

				logger.Info("Command handled", watermill.LogFields{
					"command_name": params.CommandName,
					"duration":     time.Since(start),
					"err":          err,
				})

				return err
			},
			Marshaler: cqrsMarshaler,
			Logger:    logger,
		},
	)
	if err != nil {
		panic(err)
	}

	eventBus, err := cqrs.NewEventBusWithConfig(eventsPublisher, cqrs.EventBusConfig{
		GeneratePublishTopic: func(params cqrs.GenerateEventPublishTopicParams) (string, error) {
			return generateEventsTopic(params.EventName), nil
		},

		OnPublish: func(params cqrs.OnEventSendParams) error {
			logger.Info("Publishing event", watermill.LogFields{
				"event_name": params.EventName,
			})

			params.Message.Metadata.Set("published_at", time.Now().String())

			return nil
		},

		Marshaler: cqrsMarshaler,
		Logger:    logger,
	})
	if err != nil {
		panic(err)
	}

	eventProcessor, err := cqrs.NewEventProcessorWithConfig(
		router,
		cqrs.EventProcessorConfig{
			GenerateSubscribeTopic: func(params cqrs.EventProcessorGenerateSubscribeTopicParams) (string, error) {
				return generateEventsTopic(params.EventName), nil
			},
			SubscriberConstructor: func(params cqrs.EventProcessorSubscriberConstructorParams) (message.Subscriber, error) {
				config := amqp.NewDurablePubSubConfig(
					amqpAddress,
					amqp.GenerateQueueNameTopicNameWithSuffix(params.HandlerName),
				)

				return amqp.NewSubscriber(config, logger)
			},

			OnHandle: func(params cqrs.EventProcessorOnHandleParams) error {
				start := time.Now()

				err := params.Handler.Handle(params.Message.Context(), params.Event)

				logger.Info("Event handled", watermill.LogFields{
					"event_name": params.EventName,
					"duration":   time.Since(start),
					"err":        err,
				})

				return err
			},

			Marshaler: cqrsMarshaler,
			Logger:    logger,
		},
	)
	if err != nil {
		panic(err)
	}

	err = commandProcessor.AddHandlers(
		cqrs.NewCommandHandler("BookRoomHandler", BookRoomHandler{eventBus}.Handle),
		cqrs.NewCommandHandler("OrderBeerHandler", OrderBeerHandler{eventBus}.Handle),
	)
	if err != nil {
		panic(err)
	}

	err = eventProcessor.AddHandlers(
		cqrs.NewEventHandler(
			"OrderBeerOnRoomBooked",
			OrderBeerOnRoomBooked{commandBus}.Handle,
		),
		cqrs.NewEventHandler(
			"LogBeerOrdered",
			func(ctx context.Context, event *BeerOrdered) error {
				logger.Info("Beer ordered", watermill.LogFields{
					"room_id": event.RoomId,
				})
				return nil
			},
		),
		cqrs.NewEventHandler(
			"BookingsFinancialReport",
			NewBookingsFinancialReport().Handle,
		),
	)
	if err != nil {
		panic(err)
	}

	go publishCommands(commandBus)

	if err := router.Run(context.Background()); err != nil {
		panic(err)
	}
}

func publishCommands(commandBus *cqrs.CommandBus) func() { _ = "STUB: not implemented"; return nil }
