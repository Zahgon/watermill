package cqrs

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type ProtoMarshaler struct {
	NewUUID      func() string
	GenerateName func(v interface{}) string
}

type NoProtoMessageError struct {
	v interface{}
}

func (e NoProtoMessageError) Error() string { _ = "STUB: not implemented"; return "" }

func (m ProtoMarshaler) Marshal(v interface{}) (*message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m ProtoMarshaler) newUUID() string { _ = "STUB: not implemented"; return "" }

func (ProtoMarshaler) Unmarshal(msg *message.Message, v interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m ProtoMarshaler) Name(cmdOrEvent interface{}) string { _ = "STUB: not implemented"; return "" }

func (m ProtoMarshaler) NameFromMessage(msg *message.Message) string {
	_ = "STUB: not implemented"
	return ""
}
