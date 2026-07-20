package main

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
)

func GenerateMessageMetadata(partitionKey string) *MessageMetadata {
	_ = "STUB: not implemented"
	return nil
}

type CqrsMarshalerDecorator struct {
	cqrs.ProtoMarshaler
}

const PartitionKeyMetadataField = "partition_key"

func (c CqrsMarshalerDecorator) Marshal(v interface{}) (*message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ProtoMessage interface {
	GetMetadata() *MessageMetadata
}

func GenerateKafkaPartitionKey(topic string, msg *message.Message) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
