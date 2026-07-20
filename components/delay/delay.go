package delay

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
)

type Delay struct {
	time     time.Time
	duration time.Duration
}

func (d Delay) IsZero() bool { _ = "STUB: not implemented"; return false }

func Until(delayedUntil time.Time) Delay { _ = "STUB: not implemented"; return *new(Delay) }

func For(delayedFor time.Duration) Delay { _ = "STUB: not implemented"; return *new(Delay) }

type contextKey string

var (
	delayContextKey = contextKey("delay")
)

func WithContext(ctx context.Context, delay Delay) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

const (
	DelayedUntilKey = "_watermill_delayed_until"
	DelayedForKey   = "_watermill_delayed_for"
)

func Message(msg *message.Message, delay Delay) { _ = "STUB: not implemented"; return }
