package cqrs

import (
	"context"
)

type CommandHandler interface {
	HandlerName() string

	NewCommand() any

	Handle(ctx context.Context, cmd any) error
}

type genericCommandHandler[Command any] struct {
	handleFunc  func(ctx context.Context, cmd *Command) error
	handlerName string
}

func NewCommandHandler[Command any](
	handlerName string,
	handleFunc func(ctx context.Context, cmd *Command) error,
) CommandHandler {
	_ = "STUB: not implemented"
	return *new(CommandHandler)
}

func (c genericCommandHandler[Command]) HandlerName() string { _ = "STUB: not implemented"; return "" }

func (c genericCommandHandler[Command]) NewCommand() any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (c genericCommandHandler[Command]) Handle(ctx context.Context, cmd any) error {
	_ = "STUB: not implemented"
	return nil
}
