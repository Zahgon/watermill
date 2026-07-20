package middleware

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type IgnoreErrors struct {
	ignoredErrors map[string]struct{}
}

func NewIgnoreErrors(errs []error) IgnoreErrors {
	_ = "STUB: not implemented"
	return *new(IgnoreErrors)
}

func (i IgnoreErrors) Middleware(h message.HandlerFunc) message.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(message.HandlerFunc)
}
