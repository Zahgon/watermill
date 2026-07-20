package forwarder

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type messageEnvelope struct {
	DestinationTopic string `json:"destination_topic"`

	UUID     string            `json:"uuid"`
	Payload  []byte            `json:"payload"`
	Metadata map[string]string `json:"metadata"`
}

func newMessageEnvelope(destTopic string, msg *message.Message) (*messageEnvelope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *messageEnvelope) validate() error { _ = "STUB: not implemented"; return nil }

func wrapMessageInEnvelope(
	destinationTopic string,
	msg *message.Message,
	marshaler Marshaler,
) (*message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unwrapMessageFromEnvelope(
	msg *message.Message,
	marshaler Marshaler,
) (destinationTopic string, unwrappedMsg *message.Message, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
