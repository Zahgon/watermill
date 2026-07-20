package middleware

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

func Duplicator(h message.HandlerFunc) message.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(message.HandlerFunc)
}
