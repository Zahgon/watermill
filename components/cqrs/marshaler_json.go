package cqrs

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type JSONMarshaler struct {
	NewUUID      func() string
	GenerateName func(v interface{}) string
}

func (m JSONMarshaler) Marshal(v interface{}) (*message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m JSONMarshaler) newUUID() string { _ = "STUB: not implemented"; return "" }

func (JSONMarshaler) Unmarshal(msg *message.Message, v interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m JSONMarshaler) Name(cmdOrEvent interface{}) string { _ = "STUB: not implemented"; return "" }

func (m JSONMarshaler) NameFromMessage(msg *message.Message) string {
	_ = "STUB: not implemented"
	return ""
}
