package middleware

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type RecoveredPanicError struct {
	V          interface{}
	Stacktrace string
}

func (p RecoveredPanicError) Error() string { _ = "STUB: not implemented"; return "" }

func Recoverer(h message.HandlerFunc) message.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(message.HandlerFunc)
}
