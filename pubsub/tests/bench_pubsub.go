package tests

import (
	"testing"

	"github.com/ThreeDotsLabs/watermill/message"
)

type BenchmarkPubSubConstructor func(n int) (message.Publisher, message.Subscriber)

func BenchSubscriber(b *testing.B, pubSubConstructor BenchmarkPubSubConstructor) {
	_ = "STUB: not implemented"
	return
}
