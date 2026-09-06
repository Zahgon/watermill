package cqrs

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type CommandEventMarshaler interface {
	Marshal(v interface{}) (*message.Message, error)

	Unmarshal(msg *message.Message, v interface{}) (err error)

	Name(v interface{}) string

	NameFromMessage(msg *message.Message) string
}

type CommandEventMarshalerDecorator struct {
	CommandEventMarshaler

	DecorateFunc func(v any, msg *message.Message) error
}

func (c CommandEventMarshalerDecorator) Marshal(v any) (*message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
