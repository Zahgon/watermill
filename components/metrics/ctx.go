package metrics

import "context"

type contextValue int

const (
	publishObserved contextValue = iota
	subscribeObserved
)

func setPublishObservedToCtx(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func publishAlreadyObserved(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func setSubscribeObservedToCtx(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func subscribeAlreadyObserved(ctx context.Context) bool { _ = "STUB: not implemented"; return false }
