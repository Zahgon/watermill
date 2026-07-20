package cqrs

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type ProtobufMarshaler struct {
	NewUUID      func() string
	GenerateName func(v interface{}) string

	DisableStdProtoFallback bool
}

func (m ProtobufMarshaler) Marshal(v interface{}) (msg *message.Message, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m ProtobufMarshaler) newUUID() string { _ = "STUB: not implemented"; return "" }

func (m ProtobufMarshaler) Unmarshal(msg *message.Message, v interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m ProtobufMarshaler) ToProtoMarshaler() ProtoMarshaler {
	_ = "STUB: not implemented"
	return *new(ProtoMarshaler)
}

func (m ProtobufMarshaler) Name(cmdOrEvent interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

func (m ProtobufMarshaler) NameFromMessage(msg *message.Message) string {
	_ = "STUB: not implemented"
	return ""
}
