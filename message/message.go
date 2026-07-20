package message

import (
	"context"
	"sync"
)

var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

type Payload []byte

type Message struct {
	UUID string

	Metadata Metadata

	Payload Payload

	ack chan struct{}

	noAck chan struct{}

	ackMutex    sync.Mutex
	ackSentType ackType

	ctx context.Context
}

func NewMessage(uuid string, payload Payload) *Message { _ = "STUB: not implemented"; return nil }

func NewMessageWithContext(ctx context.Context, uuid string, payload Payload) *Message {
	_ = "STUB: not implemented"
	return nil
}

type ackType int

const (
	noAckSent ackType = iota
	ack
	nack
)

func (m *Message) Equals(toCompare *Message) bool { _ = "STUB: not implemented"; return false }

func (m *Message) Ack() bool { _ = "STUB: not implemented"; return false }

func (m *Message) Nack() bool { _ = "STUB: not implemented"; return false }

func (m *Message) Acked() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (m *Message) Nacked() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (m *Message) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (m *Message) SetContext(ctx context.Context) { _ = "STUB: not implemented"; return }

func (m *Message) Copy() *Message { _ = "STUB: not implemented"; return nil }

func (m *Message) CopyWithContext() *Message { _ = "STUB: not implemented"; return nil }
