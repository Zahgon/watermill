package metrics

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	subscriberLabelKeys = []string{
		labelKeyHandlerName,
		labelKeySubscriberName,
	}
)

type SubscriberPrometheusMetricsDecorator struct {
	message.Subscriber
	subscriberName                  string
	subscriberMessagesReceivedTotal *prometheus.CounterVec
	closing                         chan struct{}
	additionalLabels                []MetricLabel
}

func (s SubscriberPrometheusMetricsDecorator) recordMetrics(msg *message.Message) {
	_ = "STUB: not implemented"
	return
}
