package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

func CreateRegistryAndServeHTTP(addr string) (registry *prometheus.Registry, cancel func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ServeHTTP(addr string, registry *prometheus.Registry) (cancel func()) {
	_ = "STUB: not implemented"
	return nil
}
