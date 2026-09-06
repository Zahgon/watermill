package subscriber

import (
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
)

func BulkRead(messagesCh <-chan *message.Message, limit int, timeout time.Duration) (receivedMessages message.Messages, all bool) {
	_ = "STUB: not implemented"
	return *new(message.Messages), false
}

func BulkReadWithDeduplication(messagesCh <-chan *message.Message, limit int, timeout time.Duration) (receivedMessages message.Messages, all bool) {
	_ = "STUB: not implemented"
	return *new(message.Messages), false
}
