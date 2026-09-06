package middleware

import (
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
)

func Timeout(timeout time.Duration) func(message.HandlerFunc) message.HandlerFunc {
	_ = "STUB: not implemented"
	return nil
}
