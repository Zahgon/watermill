package tests

import (
	"context"
	"testing"

	"github.com/ThreeDotsLabs/watermill/message"
)

func difference(a, b []string) []string { _ = "STUB: not implemented"; return nil }

func MissingMessages(expected message.Messages, received message.Messages) []string {
	_ = "STUB: not implemented"
	return nil
}

func AssertAllMessagesReceived(t *testing.T, sent message.Messages, received message.Messages) bool {
	_ = "STUB: not implemented"
	return false
}

func AssertMessagesPayloads(
	t *testing.T,
	expectedPayloads map[string][]byte,
	received []*message.Message,
) bool {
	_ = "STUB: not implemented"
	return false
}

func AssertMessagesMetadata(t *testing.T, key string, expectedValues map[string]string, received []*message.Message) bool {
	_ = "STUB: not implemented"
	return false
}

func AssertAllMessagesHaveSameContext(t *testing.T, contextKeyString string, expectedValues map[string]context.Context, received []*message.Message) {
	_ = "STUB: not implemented"
	return
}
