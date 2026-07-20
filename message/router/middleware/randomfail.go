package middleware

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

func shouldFail(probability float32) bool { _ = "STUB: not implemented"; return false }

func RandomFail(errorProbability float32) message.HandlerMiddleware {
	_ = "STUB: not implemented"
	return *new(message.HandlerMiddleware)
}

func RandomPanic(panicProbability float32) message.HandlerMiddleware {
	_ = "STUB: not implemented"
	return *new(message.HandlerMiddleware)
}
