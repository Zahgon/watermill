package cqrs

import (
	"context"
)

type EventHandler interface {
	HandlerName() string

	NewEvent() any

	Handle(ctx context.Context, event any) error
}

type genericEventHandler[T any] struct {
	handleFunc  func(ctx context.Context, event *T) error
	handlerName string
}

func NewEventHandler[T any](
	handlerName string,
	handleFunc func(ctx context.Context, event *T) error,
) EventHandler {
	_ = "STUB: not implemented"
	return *new(EventHandler)
}

func (c genericEventHandler[T]) HandlerName() string { _ = "STUB: not implemented"; return "" }

func (c genericEventHandler[T]) NewEvent() any { _ = "STUB: not implemented"; return *new(any) }

func (c genericEventHandler[T]) Handle(ctx context.Context, e any) error {
	_ = "STUB: not implemented"
	return nil
}

type GroupEventHandler interface {
	NewEvent() interface{}
	Handle(ctx context.Context, event interface{}) error
}

func NewGroupEventHandler[T any](handleFunc func(ctx context.Context, event *T) error) GroupEventHandler {
	_ = "STUB: not implemented"
	return *new(GroupEventHandler)
}
