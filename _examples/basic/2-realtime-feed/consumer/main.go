package main

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
)

var (
	marshaler = kafka.DefaultMarshaler{}
	brokers   = []string{"kafka:9092"}
)

func main() {
	logger := watermill.NewStdLogger(false, false)
	logger.Info("Starting the consumer", nil)

	pub, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   brokers,
			Marshaler: marshaler,
		},
		logger,
	)
	if err != nil {
		panic(err)
	}

	r, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	retryMiddleware := middleware.Retry{
		MaxRetries:      1,
		InitialInterval: time.Millisecond * 10,
	}

	poisonQueue, err := middleware.PoisonQueue(pub, "poison_queue")
	if err != nil {
		panic(err)
	}

	r.AddMiddleware(

		middleware.Recoverer,

		middleware.NewThrottle(10, time.Second).Middleware,

		poisonQueue,

		retryMiddleware.Middleware,

		middleware.CorrelationID,

		middleware.RandomFail(0.01),
		middleware.RandomPanic(0.01),
	)

	r.AddPlugin(plugin.SignalsHandler)

	r.AddHandler(
		"posts_counter",
		"posts_published",
		createSubscriber("posts_counter", logger),
		"posts_count",
		pub,
		PostsCounter{memoryCountStorage{new(int64)}}.Count,
	)

	r.AddConsumerHandler(
		"feed_generator",
		"posts_published",
		createSubscriber("feed_generator", logger),
		FeedGenerator{printFeedStorage{}}.UpdateFeed,
	)

	if err = r.Run(context.Background()); err != nil {
		panic(err)
	}
}

func createSubscriber(consumerGroup string, logger watermill.LoggerAdapter) message.Subscriber {
	_ = "STUB: not implemented"
	return *new(message.Subscriber)
}

type postsCountUpdated struct {
	NewCount int64 `json:"new_count"`
}

type countStorage interface {
	CountAdd() (int64, error)
	Count() (int64, error)
}

type memoryCountStorage struct {
	count *int64
}

func (m memoryCountStorage) CountAdd() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (m memoryCountStorage) Count() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

type PostsCounter struct {
	countStorage countStorage
}

func (p PostsCounter) Count(msg *message.Message) ([]*message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type postAdded struct {
	OccurredOn time.Time `json:"occurred_on"`
	Author     string    `json:"author"`
	Title      string    `json:"title"`
}

type feedStorage interface {
	AddToFeed(title, author string, time time.Time) error
}

type printFeedStorage struct{}

func (printFeedStorage) AddToFeed(title, author string, time time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

type FeedGenerator struct {
	feedStorage feedStorage
}

func (f FeedGenerator) UpdateFeed(message *message.Message) error {
	_ = "STUB: not implemented"
	return nil
}
