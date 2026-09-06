package middleware

import (
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type RetryParams struct {
	Err error

	RetryNum int

	Delay time.Duration
}

type RetriesExhaustedParams struct {
	Err error

	RetryNum int
}

type Retry struct {
	MaxRetries int

	InitialInterval time.Duration

	MaxInterval time.Duration

	Multiplier float64

	MaxElapsedTime time.Duration

	RandomizationFactor float64

	OnRetryHook func(retryNum int, delay time.Duration)

	OnRetriesExhausted func(params RetriesExhaustedParams)

	ShouldRetry func(params RetryParams) bool

	ResetContextOnRetry bool

	Logger watermill.LoggerAdapter
}

func (r Retry) Middleware(h message.HandlerFunc) message.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(message.HandlerFunc)
}
