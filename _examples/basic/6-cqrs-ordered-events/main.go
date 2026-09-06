package main

import (
	"context"
	"log/slog"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	logger := watermill.NewSlogLoggerWithLevelMapping(nil, map[slog.Level]slog.Level{
		slog.LevelInfo: slog.LevelDebug,
	})

	cqrsMarshaler := CqrsMarshalerDecorator{
		cqrs.ProtoMarshaler{

			GenerateName: cqrs.StructName,
		},
	}

	watermillLogger := watermill.NewSlogLoggerWithLevelMapping(
		slog.With("watermill", true),
		map[slog.Level]slog.Level{
			slog.LevelInfo: slog.LevelDebug,
		},
	)

	kafkaMarshaler := kafka.NewWithPartitioningMarshaler(GenerateKafkaPartitionKey)

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{"kafka:9092"},
			Marshaler: kafkaMarshaler,
		},
		watermillLogger,
	)
	if err != nil {
		panic(err)
	}

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddMiddleware(middleware.Recoverer)
	router.AddMiddleware(func(h message.HandlerFunc) message.HandlerFunc {
		return func(msg *message.Message) ([]*message.Message, error) {
			slog.Debug("Received message", "metadata", msg.Metadata)
			return h(msg)
		}
	})

	commandBus, err := cqrs.NewCommandBusWithConfig(publisher, cqrs.CommandBusConfig{
		GeneratePublishTopic: func(params cqrs.CommandBusGeneratePublishTopicParams) (string, error) {

			return "commands", nil
		},
		Marshaler: cqrsMarshaler,
		Logger:    logger,
	})
	if err != nil {
		panic(err)
	}

	eventBus, err := cqrs.NewEventBusWithConfig(publisher, cqrs.EventBusConfig{
		GeneratePublishTopic: func(params cqrs.GenerateEventPublishTopicParams) (string, error) {

			return "events", nil
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
				return "commands", nil
			},
			SubscriberConstructor: func(params cqrs.CommandProcessorSubscriberConstructorParams) (message.Subscriber, error) {
				return kafka.NewSubscriber(
					kafka.SubscriberConfig{
						Brokers:       []string{"kafka:9092"},
						ConsumerGroup: params.HandlerName,
						Unmarshaler:   kafkaMarshaler,
					},
					watermillLogger,
				)
			},
			Marshaler: cqrsMarshaler,
			Logger:    logger,
		},
	)
	if err != nil {
		panic(err)
	}

	eventProcessor, err := cqrs.NewEventGroupProcessorWithConfig(
		router,
		cqrs.EventGroupProcessorConfig{
			GenerateSubscribeTopic: func(params cqrs.EventGroupProcessorGenerateSubscribeTopicParams) (string, error) {
				return "events", nil
			},
			SubscriberConstructor: func(params cqrs.EventGroupProcessorSubscriberConstructorParams) (message.Subscriber, error) {
				return kafka.NewSubscriber(
					kafka.SubscriberConfig{
						Brokers:       []string{"kafka:9092"},
						ConsumerGroup: params.EventGroupName,
						Unmarshaler:   kafkaMarshaler,
					},
					watermillLogger,
				)
			},
			Marshaler: cqrsMarshaler,
			Logger:    logger,
		},
	)
	if err != nil {
		panic(err)
	}

	err = commandProcessor.AddHandlers(
		cqrs.NewCommandHandler("SubscribeHandler", SubscribeHandler{eventBus}.Handle),
		cqrs.NewCommandHandler("UnsubscribeHandler", UnsubscribeHandler{eventBus}.Handle),
		cqrs.NewCommandHandler("UpdateEmailHandler", UpdateEmailHandler{eventBus}.Handle),
	)
	if err != nil {
		panic(err)
	}

	subscribersReadModel := NewSubscriberReadModel()

	err = eventProcessor.AddHandlersGroup(
		"SubscriberReadModel",
		cqrs.NewGroupEventHandler(subscribersReadModel.OnSubscribed),
		cqrs.NewGroupEventHandler(subscribersReadModel.OnUnsubscribed),
		cqrs.NewGroupEventHandler(subscribersReadModel.OnEmailUpdated),
	)
	if err != nil {
		panic(err)
	}

	activityReadModel := NewActivityTimelineModel()

	err = eventProcessor.AddHandlersGroup(
		"ActivityTimelineReadModel",
		cqrs.NewGroupEventHandler(activityReadModel.OnSubscribed),
		cqrs.NewGroupEventHandler(activityReadModel.OnUnsubscribed),
		cqrs.NewGroupEventHandler(activityReadModel.OnEmailUpdated),
	)
	if err != nil {
		panic(err)
	}

	slog.Info("Starting service")

	go simulateTraffic(commandBus)

	if err := router.Run(context.Background()); err != nil {
		panic(err)
	}
}

func simulateTraffic(commandBus *cqrs.CommandBus) { _ = "STUB: not implemented"; return }

type SubscribeHandler struct {
	eventBus *cqrs.EventBus
}

func (h SubscribeHandler) Handle(ctx context.Context, cmd *Subscribe) error {
	_ = "STUB: not implemented"
	return nil
}

type UnsubscribeHandler struct {
	eventBus *cqrs.EventBus
}

func (h UnsubscribeHandler) Handle(ctx context.Context, cmd *Unsubscribe) error {
	_ = "STUB: not implemented"
	return nil
}

type UpdateEmailHandler struct {
	eventBus *cqrs.EventBus
}

func (h UpdateEmailHandler) Handle(ctx context.Context, cmd *UpdateEmail) error {
	_ = "STUB: not implemented"
	return nil
}
