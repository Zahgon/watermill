package middleware

import (
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
)

type Throttle struct {
	ticker *time.Ticker
}

func NewThrottle(count int64, duration time.Duration) *Throttle {
	_ = "STUB: not implemented"
	return nil
}

func (t Throttle) Middleware(h message.HandlerFunc) message.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(message.HandlerFunc)
}
