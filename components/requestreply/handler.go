package requestreply

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
)

func NewCommandHandler[Command any](
	handlerName string,
	backend Backend[struct{}],
	handleFunc func(ctx context.Context, cmd *Command) error,
) cqrs.CommandHandler {
	_ = "STUB: not implemented"
	return *new(cqrs.CommandHandler)
}

func NewCommandHandlerWithResult[Command any, Result any](
	handlerName string,
	backend Backend[Result],
	handleFunc func(ctx context.Context, cmd *Command) (Result, error),
) cqrs.CommandHandler {
	_ = "STUB: not implemented"
	return *new(cqrs.CommandHandler)
}

func originalCommandMsgFromCtx(ctx context.Context) (*message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
