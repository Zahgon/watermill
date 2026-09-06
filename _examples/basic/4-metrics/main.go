package main

import (
	"context"
	"flag"
	"math/rand"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/metrics"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

var (
	metricsAddr  = flag.String("metrics", ":8081", "The address that will expose /metrics for Prometheus")
	handlerDelay = flag.Float64("delay", 0, "The stdev of normal distribution of delay in handler (in seconds), to simulate load")

	logger = watermill.NewStdLogger(true, true)
	random = rand.New(rand.NewSource(time.Now().Unix()))
)

func delay() { _ = "STUB: not implemented"; return }

func handler(msg *message.Message) ([]*message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func consumeMessages(subscriber message.Subscriber) { _ = "STUB: not implemented"; return }

func produceMessages(routerClosed chan struct{}, publisher message.Publisher) {
	_ = "STUB: not implemented"
	return
}

func main() {
	flag.Parse()

	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)

	router, err := message.NewRouter(
		message.RouterConfig{},
		logger,
	)
	if err != nil {
		panic(err)
	}

	prometheusRegistry, closeMetricsServer := metrics.CreateRegistryAndServeHTTP(*metricsAddr)
	defer closeMetricsServer()

	metricsBuilder := metrics.NewPrometheusMetricsBuilder(prometheusRegistry, "", "")
	metricsBuilder.AddPrometheusRouterMetrics(router)

	router.AddMiddleware(
		middleware.Recoverer,
		middleware.RandomFail(0.1),
		middleware.RandomPanic(0.1),
	)
	router.AddPlugin(plugin.SignalsHandler)

	router.AddHandler(
		"metrics-example",
		"sub_topic",
		pubSub,
		"pub_topic",
		pubSub,
		handler,
	)

	pub := randomFailPublisherDecorator{pubSub, 0.1}

	subWithMetrics, err := metricsBuilder.DecorateSubscriber(pubSub)
	if err != nil {
		panic(err)
	}
	pubWithMetrics, err := metricsBuilder.DecoratePublisher(pub)
	if err != nil {
		panic(err)
	}

	routerClosed := make(chan struct{})
	go produceMessages(routerClosed, pubWithMetrics)
	go consumeMessages(subWithMetrics)

	_ = router.Run(context.Background())
	close(routerClosed)
}

type randomFailPublisherDecorator struct {
	message.Publisher
	failProbability float64
}

func (r randomFailPublisherDecorator) Publish(topic string, messages ...*message.Message) error {
	_ = "STUB: not implemented"
	return nil
}
