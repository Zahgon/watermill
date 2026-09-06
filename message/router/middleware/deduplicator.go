package middleware

import (
	"context"
	"sync"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
)

const MessageHasherReadLimitMinimum = 64

type ExpiringKeyRepository interface {
	IsDuplicate(ctx context.Context, key string) (ok bool, err error)
}

type MessageHasher func(*message.Message) (string, error)

type Deduplicator struct {
	KeyFactory MessageHasher
	Repository ExpiringKeyRepository
	Timeout    time.Duration
}

func (d *Deduplicator) IsDuplicate(m *message.Message) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func applyDefaultsToDeduplicator(d *Deduplicator) *Deduplicator {
	_ = "STUB: not implemented"
	return nil
}

func (d *Deduplicator) Middleware(h message.HandlerFunc) message.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(message.HandlerFunc)
}

type mapExpiringKeyRepository struct {
	window time.Duration
	mu     *sync.Mutex
	tags   map[string]time.Time
}

func NewMapExpiringKeyRepository(window time.Duration) (ExpiringKeyRepository, error) {
	_ = "STUB: not implemented"
	return *new(ExpiringKeyRepository), nil
}

func (kr *mapExpiringKeyRepository) IsDuplicate(
	ctx context.Context,
	key string,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (kr *mapExpiringKeyRepository) cleanOutLoop(ctx context.Context, ticker *time.Ticker) {
	_ = "STUB: not implemented"
	return
}

func (kr *mapExpiringKeyRepository) cleanOut(tagsBefore time.Time) {
	_ = "STUB: not implemented"
	return
}

func (kr *mapExpiringKeyRepository) Len() (count int) { _ = "STUB: not implemented"; return 0 }

func NewMessageHasherAdler32(readLimit int64) MessageHasher {
	_ = "STUB: not implemented"
	return *new(MessageHasher)
}

func NewMessageHasherSHA256(readLimit int64) MessageHasher {
	_ = "STUB: not implemented"
	return *new(MessageHasher)
}

func NewMessageHasherFromMetadataField(field string) MessageHasher {
	_ = "STUB: not implemented"
	return *new(MessageHasher)
}

type deduplicatingPublisherDecorator struct {
	message.Publisher
	deduplicator *Deduplicator
}

func (d *deduplicatingPublisherDecorator) Publish(
	topic string,
	messages ...*message.Message,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *Deduplicator) PublisherDecorator() message.PublisherDecorator {
	_ = "STUB: not implemented"
	return *new(message.PublisherDecorator)
}
