package middleware

import (
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
)

type DelayOnError struct {
	InitialInterval time.Duration

	MaxInterval time.Duration

	Multiplier float64
}

func (d *DelayOnError) Middleware(h message.HandlerFunc) message.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(message.HandlerFunc)
}

func (d *DelayOnError) applyDelay(msg *message.Message) { _ = "STUB: not implemented"; return }
