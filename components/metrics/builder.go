package metrics

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusMetricsBuilderConfig struct {
	Namespace        string
	Subsystem        string
	AdditionalLabels []MetricLabel
}

func NewPrometheusMetricsBuilderWithConfig(prometheusRegistry prometheus.Registerer, config PrometheusMetricsBuilderConfig) PrometheusMetricsBuilder {
	_ = "STUB: not implemented"
	return *new(PrometheusMetricsBuilder)
}

func NewPrometheusMetricsBuilder(prometheusRegistry prometheus.Registerer, namespace string, subsystem string) PrometheusMetricsBuilder {
	_ = "STUB: not implemented"
	return *new(PrometheusMetricsBuilder)
}

type PrometheusMetricsBuilder struct {
	PrometheusRegistry prometheus.Registerer

	Namespace string
	Subsystem string

	PublishBuckets []float64

	HandlerBuckets []float64

	additionalLabels []MetricLabel
}

func (b PrometheusMetricsBuilder) AddPrometheusRouterMetrics(r *message.Router) {
	_ = "STUB: not implemented"
	return
}

func (b PrometheusMetricsBuilder) DecoratePublisher(pub message.Publisher) (message.Publisher, error) {
	_ = "STUB: not implemented"
	return *new(message.Publisher), nil
}

func (b PrometheusMetricsBuilder) DecorateSubscriber(sub message.Subscriber) (message.Subscriber, error) {
	_ = "STUB: not implemented"
	return *new(message.Subscriber), nil
}

func (b PrometheusMetricsBuilder) register(c prometheus.Collector) (prometheus.Collector, error) {
	_ = "STUB: not implemented"
	return *new(prometheus.Collector), nil
}

func (b PrometheusMetricsBuilder) registerCounterVec(c *prometheus.CounterVec) (*prometheus.CounterVec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b PrometheusMetricsBuilder) registerHistogramVec(h *prometheus.HistogramVec) (*prometheus.HistogramVec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
