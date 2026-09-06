package cli

import (
	"time"
)

type Message struct {
	ID       string
	UUID     string
	Payload  string
	Metadata map[string]string

	OriginalTopic string
	DelayedUntil  string
	DelayedFor    string
	RequeueIn     time.Duration
}

func NewMessage(id string, uuid string, payload string, metadata map[string]string) (Message, error) {
	_ = "STUB: not implemented"
	return *new(Message), nil
}
