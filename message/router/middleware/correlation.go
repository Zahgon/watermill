package middleware

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

const CorrelationIDMetadataKey = "correlation_id"

func SetCorrelationID(id string, msg *message.Message) { _ = "STUB: not implemented"; return }

func MessageCorrelationID(message *message.Message) string { _ = "STUB: not implemented"; return "" }

func CorrelationID(h message.HandlerFunc) message.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(message.HandlerFunc)
}
