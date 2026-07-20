package metrics

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	publisherLabelKeys = []string{
		labelKeyHandlerName,
		labelKeyPublisherName,
		labelSuccess,
	}
)

type PublisherPrometheusMetricsDecorator struct {
	pub                message.Publisher
	publisherName      string
	publishTimeSeconds *prometheus.HistogramVec
	additionalLabels   []MetricLabel
}

func (m PublisherPrometheusMetricsDecorator) Publish(topic string, messages ...*message.Message) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m PublisherPrometheusMetricsDecorator) Close() error { _ = "STUB: not implemented"; return nil }
