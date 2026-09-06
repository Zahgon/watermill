package metrics

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	labelKeyHandlerName    = "handler_name"
	labelKeyPublisherName  = "publisher_name"
	labelKeySubscriberName = "subscriber_name"
	labelSuccess           = "success"
	labelAcked             = "acked"

	labelValueNoHandler = "<no handler>"
)

var (
	labelGetters = map[string]func(context.Context) string{
		labelKeyHandlerName:    message.HandlerNameFromCtx,
		labelKeyPublisherName:  message.PublisherNameFromCtx,
		labelKeySubscriberName: message.SubscriberNameFromCtx,
	}
)

func labelsFromCtx(ctx context.Context, labels ...string) prometheus.Labels {
	_ = "STUB: not implemented"
	return *new(prometheus.Labels)
}

type LabelComputeValueFn func(msgCtx context.Context) string

type MetricLabel struct {
	Label          string
	ComputeValueFn LabelComputeValueFn
}

func toLabelsSlice(baseLabels []string, customs []MetricLabel) []string {
	_ = "STUB: not implemented"
	return nil
}
