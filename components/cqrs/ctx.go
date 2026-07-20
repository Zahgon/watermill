package cqrs

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

type ctxKey string

const (
	originalMessage ctxKey = "original_message"
)

func OriginalMessageFromCtx(ctx context.Context) *message.Message {
	_ = "STUB: not implemented"
	return nil
}

func CtxWithOriginalMessage(ctx context.Context, msg *message.Message) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
