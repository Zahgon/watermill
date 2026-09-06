package middleware

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
)

var ErrInvalidPoisonQueueTopic = errors.New("invalid poison queue topic")

const (
	ReasonForPoisonedKey  = "reason_poisoned"
	PoisonedTopicKey      = "topic_poisoned"
	PoisonedHandlerKey    = "handler_poisoned"
	PoisonedSubscriberKey = "subscriber_poisoned"
)

type poisonQueue struct {
	topic string
	pub   message.Publisher

	shouldGoToPoisonQueue func(err error) bool
}

func PoisonQueue(pub message.Publisher, topic string) (message.HandlerMiddleware, error) {
	_ = "STUB: not implemented"
	return *new(message.HandlerMiddleware), nil
}

func PoisonQueueWithFilter(pub message.Publisher, topic string, shouldGoToPoisonQueue func(err error) bool) (message.HandlerMiddleware, error) {
	_ = "STUB: not implemented"
	return *new(message.HandlerMiddleware), nil
}

func (pq poisonQueue) publishPoisonMessage(msg *message.Message, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (pq poisonQueue) Middleware(h message.HandlerFunc) message.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(message.HandlerFunc)
}
