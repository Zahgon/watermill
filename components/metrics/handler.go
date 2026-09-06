package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ThreeDotsLabs/watermill/message"
)

var (
	handlerLabelKeys = []string{
		labelKeyHandlerName,
		labelSuccess,
	}

	defaultHandlerExecutionTimeBuckets = []float64{
		0.0005,
		0.001,
		0.0025,
		0.005,
		0.01,
		0.025,
		0.05,
		0.1,
		0.25,
		0.5,
		1,
	}
)

type HandlerPrometheusMetricsMiddleware struct {
	handlerExecutionTimeSeconds *prometheus.HistogramVec
	additionalLabels            []MetricLabel
}

func (m HandlerPrometheusMetricsMiddleware) Middleware(h message.HandlerFunc) message.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(message.HandlerFunc)
}

func (b PrometheusMetricsBuilder) NewRouterMiddleware() HandlerPrometheusMetricsMiddleware {
	_ = "STUB: not implemented"
	return *new(HandlerPrometheusMetricsMiddleware)
}
