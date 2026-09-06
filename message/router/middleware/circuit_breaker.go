package middleware

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/sony/gobreaker"
)

type CircuitBreaker struct {
	cb *gobreaker.CircuitBreaker
}

func NewCircuitBreaker(settings gobreaker.Settings) CircuitBreaker {
	_ = "STUB: not implemented"
	return *new(CircuitBreaker)
}

func (c CircuitBreaker) Middleware(h message.HandlerFunc) message.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(message.HandlerFunc)
}
