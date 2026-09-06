package common

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

const UpdatesTopic = "updates"

func NotifyMiddleware(pub message.Publisher, serviceName string) func(message.HandlerFunc) message.HandlerFunc {
	_ = "STUB: not implemented"
	return nil
}
