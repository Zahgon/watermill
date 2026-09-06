package main

import (
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

var (
	brokers = []string{"kafka:9092"}

	messagesPerSecond = 100
	numWorkers        = 20
)

func main() {
	logger := watermill.NewStdLogger(false, false)
	logger.Info("Starting the producer", watermill.LogFields{})

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   brokers,
			Marshaler: kafka.DefaultMarshaler{},
		},
		logger,
	)
	if err != nil {
		panic(err)
	}
	defer publisher.Close()

	closeCh := make(chan struct{})
	workersGroup := &sync.WaitGroup{}
	workersGroup.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(publisher, workersGroup, closeCh)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	close(closeCh)

	workersGroup.Wait()

	logger.Info("All messages published", nil)
}

func worker(publisher message.Publisher, wg *sync.WaitGroup, closeCh chan struct{}) {
	_ = "STUB: not implemented"
	return
}

type postAdded struct {
	OccurredOn time.Time `json:"occurred_on"`

	Author string `json:"author"`
	Title  string `json:"title"`

	Content string `json:"content"`
}
