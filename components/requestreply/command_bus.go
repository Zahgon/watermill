package requestreply

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

type CommandBus interface {
	SendWithModifiedMessage(ctx context.Context, cmd any, modify func(*message.Message) error) error
}

func SendWithReply[Result any](
	ctx context.Context,
	c CommandBus,
	backend Backend[Result],
	cmd any,
) (Reply[Result], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SendWithReplies[Result any](
	ctx context.Context,
	c CommandBus,
	backend Backend[Result],
	cmd any,
) (replCh <-chan Reply[Result], cancel func(), err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
