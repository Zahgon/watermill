package requestreply

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
)

type NoResult = struct{}

type Reply[Result any] struct {
	HandlerResult Result

	Error error

	NotificationMessage *message.Message
}

type Backend[Result any] interface {
	ListenForNotifications(ctx context.Context, params BackendListenForNotificationsParams) (<-chan Reply[Result], error)
	OnCommandProcessed(ctx context.Context, params BackendOnCommandProcessedParams[Result]) error
}

type BackendListenForNotificationsParams struct {
	Command     any
	OperationID OperationID
}

type BackendOnCommandProcessedParams[Result any] struct {
	Command        any
	CommandMessage *message.Message

	HandlerResult Result
	HandleErr     error
}

type OperationID string

type ReplyTimeoutError struct {
	Duration time.Duration
	Err      error
}

func (e ReplyTimeoutError) Error() string { _ = "STUB: not implemented"; return "" }

type ReplyUnmarshalError struct {
	Err error
}

func (r ReplyUnmarshalError) Error() string { _ = "STUB: not implemented"; return "" }

func (r ReplyUnmarshalError) Unwrap() error { _ = "STUB: not implemented"; return nil }

type CommandHandlerError struct {
	Err error
}

func (e CommandHandlerError) Error() string { _ = "STUB: not implemented"; return "" }

func (e CommandHandlerError) Unwrap() error { _ = "STUB: not implemented"; return nil }
